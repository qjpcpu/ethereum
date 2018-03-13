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
	"time"
)

type TransactionScanner struct {
	contractAddr common.Address
	conn         *ethclient.Client
	TotalSupply  *big.Int
	Name         string
	Symbol       string
	db           StatStorageInterface
}

type TransferPacket struct {
	BlockNumber *big.Int
	Timestamp   time.Time
	Records     []TransferRecord
}

type TransferRecord struct {
	TxHash string
	From   string
	To     string
	Amount *big.Int
}

type StatStorageInterface interface {
	RecieveRecords(TransferPacket)
}

type StatStoragePrinter struct{}

func GetScanner(rawurl string, contractAddr string, storage StatStorageInterface) (*TransactionScanner, error) {
	scanner := &TransactionScanner{}
	var err error
	for loop := true; loop; loop = false {
		if scanner.conn, err = ethclient.Dial(rawurl); err != nil {
			break
		}
		if !contracts.IsContract(scanner.conn, contractAddr) {
			err = errors.New("bad contract address")
			break
		}
		var token *erc20.Token
		token, err = erc20.NewToken(common.HexToAddress(contractAddr), scanner.conn)
		if err != nil {
			log.Errorf("instantiate contract fail:%v", err)
			break
		}
		scanner.Name, _ = token.Name(nil)
		scanner.Symbol, _ = token.Symbol(nil)
		scanner.TotalSupply, _ = token.TotalSupply(nil)
		scanner.contractAddr = common.HexToAddress(contractAddr)
		scanner.db = storage
		log.Debugf("init scanner for (%s,%s (total:%s) OK", scanner.Name, scanner.Symbol, scanner.TotalSupply.String())
	}
	return scanner, err
}

func (ts *TransactionScanner) StartScan(start_block *big.Int, limit uint64) error {
	channel := make(chan TransferPacket)
	finish := make(chan struct{})
	go func() {
		for {
			select {
			case packet := <-channel:
				ts.db.RecieveRecords(packet)
			case <-finish:
				close(finish)
				close(channel)
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
			return err
		}
		block_time := time.Unix(block.Time().Int64(), 0)
		txs := block.Transactions()
		log.Debugf("got %d transactions in block %s", len(txs), start_block.String())
		var records []TransferRecord
		for _, tx := range txs {
			toAddr := tx.To()
			if toAddr == nil || *toAddr != ts.contractAddr {
				continue
			}
			if erc20.IsTransferFunc(tx.Data()) {
				to, amount, err := erc20.DecodeTransferData(tx.Data())
				if err != nil {
					log.Errorf("decode transaction %v fail:%v", tx, err)
					return err
				}
				from := (&contracts.TransactionWithExtra{Transaction: tx}).From()
				log.Debugf("Transaction:%s From:%s To:%s Amount:%s(%s)", tx.Hash().Hex(), from.Hex(), to.Hex(), amount, ts.Symbol)
				records = append(records, TransferRecord{
					TxHash: tx.Hash().Hex(),
					From:   from.Hex(),
					To:     to.Hex(),
					Amount: amount,
				})
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

func (s StatStoragePrinter) RecieveRecords(p TransferPacket) {
	log.Infof("recieved %d records of block %v", len(p.Records), p.BlockNumber)
	for _, record := range p.Records {
		log.Infof("%s: %s ==> %s %v", record.TxHash, record.From, record.To, record.Amount.String())
	}
}
