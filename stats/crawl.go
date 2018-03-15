package stats

import (
	"context"
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/qjpcpu/ethereum/contracts"
	"github.com/qjpcpu/ethereum/contracts/erc20"
	"github.com/qjpcpu/log"
	"math/big"
	"sync"
	"time"
)

type ContractInfo struct {
	Name        string
	Symbol      string
	TotalSupply uint64
	Address     string
}

type TransactionScanner struct {
	mycontracts map[string]ContractInfo
	conn        *ethclient.Client
	listener    TxListener
	mutex       *sync.Mutex
	scanning    bool
}

type TransferPacket struct {
	BlockNumber *big.Int
	Timestamp   time.Time
	Records     []TransferRecord
}

type TransferRecord struct {
	Contract           ContractInfo
	TxHash             string
	From               string
	IsContractCreation bool
	To                 string
	Amount             *big.Int
}

type TxListener interface {
	RecieveRecords(TransferPacket)
}

type StatPrinter struct{}

func GetScanner(rawurl string, lis TxListener) (*TransactionScanner, error) {
	if conn, err := ethclient.Dial(rawurl); err != nil {
		return nil, err
	} else {
		return GetScannerByClient(conn, lis), nil
	}
}

func GetScannerByClient(conn *ethclient.Client, lis TxListener) *TransactionScanner {
	return &TransactionScanner{
		mycontracts: make(map[string]ContractInfo),
		mutex:       &sync.Mutex{},
		conn:        conn,
		listener:    lis,
	}
}

func (ts *TransactionScanner) Reset() {
	ts.mutex.Lock()
	ts.mycontracts = make(map[string]ContractInfo)
	ts.mutex.Unlock()
}

func (ts *TransactionScanner) Subscribe(contractAddrs ...string) error {
	if ts.scanning {
		return errors.New("is running")
	}
	ts.mutex.Lock()
	defer ts.mutex.Unlock()
	for _, contractAddr := range contractAddrs {
		addr := common.HexToAddress(contractAddr)
		if !contracts.IsContract(ts.conn, contractAddr) {
			return errors.New("bad contract address")
		}
		var token *erc20.Token
		token, err := erc20.NewToken(addr, ts.conn)
		if err != nil {
			log.Errorf("instantiate contract fail:%v", err)
			return err
		}
		info := ContractInfo{}
		info.Name, _ = token.Name(nil)
		info.Symbol, _ = token.Symbol(nil)
		totalSupply, _ := token.TotalSupply(nil)
		info.TotalSupply = totalSupply.Uint64()
		ts.mycontracts[addr.Hex()] = info
		log.Infof("subscribe %s %s|%s OK", contractAddr, info.Name, info.Symbol)
	}
	return nil
}

func (ts *TransactionScanner) GetSubscribes() []ContractInfo {
	var list []ContractInfo
	for _, c := range ts.mycontracts {
		list = append(list, c)
	}
	return list
}

func (ts *TransactionScanner) SubscribeContracts(contractInfos ...ContractInfo) error {
	if ts.scanning {
		return errors.New("is running")
	}
	ts.mutex.Lock()
	defer ts.mutex.Unlock()
	for _, info := range contractInfos {
		ts.mycontracts[info.Address] = info
		log.Infof("subscribe %s %s|%s OK", info.Address, info.Name, info.Symbol)
	}
	return nil
}

func (ts *TransactionScanner) StartScan(start_block *big.Int, limit uint64) error {
	if ts.scanning {
		return errors.New("is running")
	}
	if len(ts.mycontracts) == 0 {
		return errors.New("no contracts to scan")
	}
	ts.mutex.Lock()
	ts.scanning = true
	channel := make(chan TransferPacket)
	finish := make(chan struct{})
	go func() {
		for {
			select {
			case packet := <-channel:
				ts.listener.RecieveRecords(packet)
			case <-finish:
				close(finish)
				close(channel)
				return
			}
		}
	}()
	defer func() {
		ts.scanning = false
		ts.mutex.Unlock()
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
			return err
		}
		block_time := time.Unix(block.Time().Int64(), 0)
		txs := block.Transactions()
		log.Debugf("got %d transactions in block %s", len(txs), start_block.String())
		var records []TransferRecord
		for _, tx := range txs {
			txe := &contracts.TransactionWithExtra{Transaction: tx}
			//是否合约创建交易
			if txe.IsContractCreation() {
				caddr := txe.ContractAddress()
				if info, ok := ts.mycontracts[caddr.Hex()]; ok {
					records = append(records, TransferRecord{
						Contract:           info,
						IsContractCreation: true,
						TxHash:             tx.Hash().Hex(),
						From:               txe.From().Hex(),
						To:                 "",
						Amount:             new(big.Int).SetInt64(0),
					})
				}
			} else {
				toAddr := txe.To()
				info, ok := ts.mycontracts[toAddr.Hex()]
				if ok && erc20.IsTransferFunc(tx.Data()) {
					to, amount, err := erc20.DecodeTransferData(tx.Data())
					if err != nil {
						log.Errorf("decode transaction %v fail:%v", tx, err)
						return err
					}
					from := txe.From()
					log.Debugf("Transaction:%s From:%s To:%s Amount:%s(%s)", tx.Hash().Hex(), from.Hex(), to.Hex(), amount, info.Symbol)
					records = append(records, TransferRecord{
						Contract:           info,
						IsContractCreation: false,
						TxHash:             tx.Hash().Hex(),
						From:               from.Hex(),
						To:                 to.Hex(),
						Amount:             amount,
					})
				}
			}

		}
		packet := TransferPacket{
			BlockNumber: new(big.Int).Set(start_block),
			Timestamp:   block_time,
			Records:     records,
		}
		channel <- packet
	}
	return nil
}

func (s StatPrinter) RecieveRecords(p TransferPacket) {
	log.Infof("recieved %d records of block %v", len(p.Records), p.BlockNumber)
	for _, record := range p.Records {
		log.Infof("%s: %s ==> %s %v", record.TxHash, record.From, record.To, record.Amount.String())
	}
}
