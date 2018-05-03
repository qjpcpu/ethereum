package events

import (
	"context"
	"errors"
	"fmt"
	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/qjpcpu/common/redo"
	abi "github.com/qjpcpu/ethereum/mabi"
	bind "github.com/qjpcpu/ethereum/mabi/mbind"
	"math/big"
	"strings"
	"time"
)

type Event struct {
	BlockNumber uint64
	TxHash      common.Hash
	Address     common.Address
	Name        string
	Data        abi.JSONObj
}

type Progress struct {
	From uint64
	To   uint64
}

func (evt Event) String() string {
	return fmt.Sprintf(
		`block: %v,tx: %s,address: %s,event: %s,data: %s`,
		evt.BlockNumber,
		evt.TxHash.Hex(),
		evt.Address.Hex(),
		evt.Name,
		evt.Data.String(),
	)
}

type Builder struct {
	es      *eventScanner
	abi_str string
}

func NewScanBuilder() *Builder {
	return &Builder{
		es: &eventScanner{},
	}
}

func (b *Builder) SetClient(conn *ethclient.Client) *Builder {
	b.es.conn = conn
	return b
}

func (b *Builder) SetEvents(names ...string) *Builder {
	b.es.EventNames = names
	return b
}

func (b *Builder) SetContract(addr common.Address) *Builder {
	b.es.Contract = addr
	return b
}

func (b *Builder) SetGracefullExit(yes bool) *Builder {
	b.es.GracefullExit = yes
	return b
}

func (b *Builder) SetABI(abi_str string) *Builder {
	b.abi_str = abi_str
	return b
}

func (b *Builder) SetFrom(f uint64) *Builder {
	b.es.From = f
	return b
}

func (b *Builder) SetProgressChan(pc chan<- Progress) *Builder {
	b.es.ProgressChan = pc
	return b
}

func (b *Builder) BuildAndRun(dataCh chan<- Event, errChan chan<- error, intervals ...time.Duration) (*redo.Recipet, error) {
	b.es.DataChan, b.es.ErrChan = dataCh, errChan
	if b.es.DataChan == nil {
		return nil, errors.New("data channel should not be empty")
	}
	if b.abi_str == "" {
		return nil, errors.New("need ABI")
	}
	if b.es.conn == nil {
		return nil, errors.New("no eth client")
	}
	if len(b.es.EventNames) == 0 {
		return nil, errors.New("please specify events")
	}
	var err error
	b.es.bc, err = bindContract(b.abi_str, b.es.Contract, b.es.conn)
	if err != nil {
		return nil, err
	}
	var recipet *redo.Recipet
	interval := time.Second * 3
	if len(intervals) > 0 {
		interval = intervals[0]
	}
	if b.es.GracefullExit {
		recipet = redo.PerformSafe(b.es.scan, interval)
	} else {
		recipet = redo.Perform(b.es.scan, interval)
	}
	return recipet, nil
}

type eventScanner struct {
	conn          *ethclient.Client
	Contract      common.Address
	EventNames    []string
	From          uint64
	DataChan      chan<- Event
	ErrChan       chan<- error
	ProgressChan  chan<- Progress
	GracefullExit bool
	bc            *bind.BoundContract
}

func (es *eventScanner) NewestBlockNumber() (uint64, error) {
	block, err := es.conn.BlockByNumber(context.Background(), nil)
	if err != nil {
		return 0, err
	}
	return block.NumberU64() - 1, nil
}

func (es *eventScanner) sendErr(err error) {
	if es.ErrChan != nil && err != nil {
		es.ErrChan <- err
	}
}

func (es *eventScanner) sendData(evt Event) {
	if es.DataChan != nil {
		es.DataChan <- evt
	}
}

func (es *eventScanner) scan(ctx *redo.RedoCtx) {
	newest_bn, err := es.NewestBlockNumber()
	if err != nil {
		// not send this err
		if !strings.Contains(err.Error(), "got null header for uncle") {
			es.sendErr(fmt.Errorf("query newest block number fail:%v, will retry later", err))
		}
		return
	}
	if es.From == 0 {
		es.From = newest_bn
	}
	var to_bn uint64 = newest_bn
	if to_bn <= es.From {
		return
	}
	if es.From+1000 < to_bn {
		to_bn -= 1000
	}
	var topics []common.Hash = make([]common.Hash, len(es.EventNames))
	for i, t := range es.EventNames {
		topics[i] = es.bc.EventTopic(t)
	}

	fq := ethereum.FilterQuery{
		FromBlock: new(big.Int).SetUint64(es.From),
		ToBlock:   new(big.Int).SetUint64(to_bn),
		Addresses: []common.Address{},
		Topics:    [][]common.Hash{topics},
	}
	if es.Contract != (common.Address{}) {
		fq.Addresses = append(fq.Addresses, es.Contract)
	}
	if es.ProgressChan != nil {
		es.ProgressChan <- Progress{From: es.From, To: to_bn}
	}
	logs, err := es.conn.FilterLogs(context.Background(), fq)
	if err != nil {
		es.sendErr(fmt.Errorf("filter log(%v,%v) err:%v, will retry later", es.From, to_bn, err))
		return
	}
	for _, lg := range logs {
		evt := abi.NewJSONObj()
		name, err := es.bc.UnpackMatchedLog(evt, lg)
		if err != nil {
			es.sendErr(fmt.Errorf("unpack %s log in tx(%s) fail:%v,abadon", name, lg.TxHash.Hex(), err))
			continue
		}
		es.sendData(Event{
			BlockNumber: lg.BlockNumber,
			TxHash:      lg.TxHash,
			Address:     lg.Address,
			Name:        name,
			Data:        evt,
		})
	}
	if len(logs) > 0 {
		ctx.StartNextRightNow()
	}
	es.From = to_bn + 1
}

func bindContract(abi_str string, address common.Address, backend bind.ContractBackend) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(abi_str))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, backend, backend, backend), nil
}
