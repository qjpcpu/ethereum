package txview

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/qjpcpu/ethereum/swg"
	"github.com/qjpcpu/log"
	"math/big"
	"strings"
	"sync"
	"time"
)

type contractFilters struct {
	// contract => function_signature
	list map[string][]string
	// function_signature => function_declare
	sig2name map[string]string
	*sync.RWMutex
}

func newContractFilters() *contractFilters {
	return &contractFilters{
		list:     make(map[string][]string),
		sig2name: make(map[string]string),
		RWMutex:  new(sync.RWMutex),
	}
}

func (cf *contractFilters) contains(contract_addr string, func_sigs ...string) bool {
	cf.RLock()
	defer cf.RUnlock()
	sigs, ok := cf.list[strings.ToLower(contract_addr)]
	if !ok || len(func_sigs) == 0 {
		return ok
	}
	// listen to all
	if len(sigs) == 0 {
		return true
	}
	for _, f := range sigs {
		if f == func_sigs[0] {
			return true
		}
	}
	return false
}

func (cf *contractFilters) sigToName(sig string) string {
	cf.RLock()
	name := cf.sig2name[sig]
	cf.RUnlock()
	return name
}

func (cf *contractFilters) add(contract_addr string, func_names ...string) {
	if cf.contains(contract_addr) {
		return
	}
	cf.Lock()
	defer cf.Unlock()
	var sigs = make([]string, 0)
	for _, name := range func_names {
		sig := crypto.Keccak256Hash([]byte(name)).Hex()[2:10]
		cf.sig2name[sig] = name
		sigs = append(sigs, sig)
	}
	cf.list[strings.ToLower(contract_addr)] = sigs
}

// return function_name,ishit
func (cf *contractFilters) isHit(contract_addr string, txData []byte) (string, bool) {
	if len(txData) <= 4 {
		return "", false
	}
	sig := common.Bytes2Hex(txData[:4])
	if cf.contains(contract_addr, sig) {
		return cf.sigToName(sig), true
	} else {
		return "", false
	}
}

type TxsInfo struct {
	BlockNumber *big.Int
	Timestamp   time.Time
	Txs         []TxInfo
}

type TxInfo struct {
	Tx              *types.Transaction
	ContractAddress string
	Function        string
}

type TxResult struct {
	Start *big.Int
	End   *big.Int
	Error error
}

type TxScan struct {
	filters  *contractFilters
	conn     *ethclient.Client
	receiver chan<- TxsInfo
	done     chan<- TxResult
}

func GetScanner(rawurl string, data chan<- TxsInfo, done chan<- TxResult) (*TxScan, error) {
	if conn, err := ethclient.Dial(rawurl); err != nil {
		return nil, err
	} else {
		return GetScannerByClient(conn, data, done), nil
	}
}

func GetScannerByClient(conn *ethclient.Client, data chan<- TxsInfo, done chan<- TxResult) *TxScan {
	return &TxScan{
		filters:  newContractFilters(),
		conn:     conn,
		receiver: data,
		done:     done,
	}
}

func (ts *TxScan) Subscribe(contractAddr string, func_names ...string) {
	ts.filters.add(contractAddr, func_names...)
}

func (ts *TxScan) handleTx(tx *types.Transaction, channel chan<- TxInfo) error {
	//是否合约创建交易
	if to := tx.To(); to != nil {
		func_name, hit := ts.filters.isHit(to.Hex(), tx.Data())
		if !hit {
			return nil
		}
		channel <- TxInfo{
			Tx:              tx,
			ContractAddress: strings.ToLower(to.Hex()),
			Function:        func_name,
		}
	}
	return nil
}

func minPositive(a, b int) int {
	if a == 0 || b == 0 {
		return a + b
	}
	if a > b {
		return b
	} else {
		return a
	}
}

func (ts *TxScan) StartScan(start_block *big.Int, limit uint64, maxTxParserCount int) {
	channel := make(chan TxsInfo, 1000)
	finish := make(chan struct{})
	fblock, tblock := new(big.Int).Set(start_block), new(big.Int).Add(start_block, big.NewInt(-1))
	result := &TxResult{}
	go func() {
		for {
			select {
			case packet := <-channel:
				ts.receiver <- packet
			case <-finish:
				close(finish)
				close(channel)
				result.Start = fblock
				result.End = new(big.Int).Add(tblock, big.NewInt(1))
				ts.done <- *result
				return
			}
		}
	}()
	defer func() {
		finish <- struct{}{}
	}()
	end_block := new(big.Int).SetUint64(limit)
	if limit > 0 {
		end_block = end_block.Add(end_block, start_block)
	}
	ctx := context.Background()
	for ; limit == 0 || start_block.Cmp(end_block) < 0; start_block = start_block.Add(start_block, big.NewInt(1)) {
		log.Debugf("start scan block %s", start_block.String())
		block, err := ts.conn.BlockByNumber(ctx, start_block)
		if err != nil {
			log.Errorf("fail to get block %s, %v", start_block.String(), err)
			result.Error = fmt.Errorf("fail to get block %v,%v", start_block, err)
			return
		}
		block_time := time.Unix(block.Time().Int64(), 0)
		txs := block.Transactions()
		log.Debugf("got %d raw transactions in block %s", len(txs), start_block.String())
		var records []TxInfo
		wg := swg.New(minPositive(len(txs), maxTxParserCount))
		datas := make(chan TxInfo, len(txs))
		for i := range txs {
			wg.Add()
			go func(tx *types.Transaction) {
				defer wg.Done()
				ts.handleTx(tx, datas)
			}(txs[i])
		}
		wg.Wait()
	LOOP:
		for {
			select {
			case record := <-datas:
				records = append(records, record)
			default:
				break LOOP
			}
		}
		close(datas)
		packet := TxsInfo{
			BlockNumber: new(big.Int).Set(start_block),
			Timestamp:   block_time,
			Txs:         records,
		}
		channel <- packet
		tblock.Set(start_block)
	}
}
