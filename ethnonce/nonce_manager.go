package ethnonce

import (
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/qjpcpu/log"
	"math/big"
	"strings"
	"time"
)

var (
	ErrNotInitAddress = errors.New("not initailize")
	ErrOtherHoldNonce = errors.New("others hold the nonce")
)

type ManagerCreator interface {
	SetEthClient(*ethclient.Client) ManagerCreator
	Build() *NonceManager
}

type NonceManager struct {
	Impl NonceManagerLowlevel
}

type NonceManagerLowlevel interface {
	PeekNonce(common.Address) uint64
	GiveNonce(common.Address) (uint64, error)
	SyncNonce(common.Address) (uint64, error)
	CommitNonce(common.Address, uint64, bool) error
	Close() error
}

func (n *NonceManager) PeekNonce(addr common.Address) uint64 {
	return n.Impl.PeekNonce(addr)
}

func (n *NonceManager) GiveNonce(addr common.Address) (uint64, error) {
	return n.Impl.GiveNonce(addr)
}

func (n *NonceManager) SyncNonce(addr common.Address) (uint64, error) {
	return n.Impl.SyncNonce(addr)
}

func (n *NonceManager) CommitNonce(addr common.Address, nonce_number uint64, success bool) error {
	return n.Impl.CommitNonce(addr, nonce_number, success)
}

func (n *NonceManager) Close() error {
	return n.Impl.Close()
}

func (n *NonceManager) MustGiveNonce(addr common.Address) (uint64, error) {
	var code uint64
	var err error
	for i := 0; i < 600; i++ {
		code, err = n.Impl.GiveNonce(addr)
		if err == nil || err == ErrNotInitAddress {
			break
		}
		// err == ErrOtherHoldNonce
		<-time.After(100 * time.Millisecond)
	}
	return code, err
}

func (n *NonceManager) GiveNonceForTx(addr common.Address, txJob func(nonce uint64) (*types.Transaction, error)) (*types.Transaction, error) {
	nonce, err := n.MustGiveNonce(addr)
	if err != nil {
		return nil, err
	}
	if tx, err := txJob(nonce); err != nil {
		n.Impl.CommitNonce(addr, nonce, false)
		if strings.Contains(err.Error(), "nonce too low") || strings.Contains(err.Error(), "nonce too high") {
			new_nonce, _ := n.Impl.SyncNonce(addr)
			log.Debugf("nonce:%d of %s is [%v], auto sync to %d", nonce, addr.Hex(), err, new_nonce)
		}
		return nil, err
	} else {
		n.Impl.CommitNonce(addr, nonce, true)
		return tx, nil
	}
}

type TxRunner interface {
	Nonce() *big.Int
	Commit(*types.Transaction, error)
}

type txRunnerImpl struct {
	nonce uint64
	tx    *types.Transaction
	err   error
}

func newTxRunner(nonce uint64) *txRunnerImpl {
	return &txRunnerImpl{
		nonce: nonce,
	}
}

func (tr *txRunnerImpl) Nonce() *big.Int {
	return new(big.Int).SetUint64(tr.nonce)
}

func (tr *txRunnerImpl) Commit(tx *types.Transaction, err error) {
	tr.tx = tx
	tr.err = err
}

func (n *NonceManager) GiveNonceForTxV2(addr common.Address, txJob func(TxRunner)) (*types.Transaction, error) {
	nonce, err := n.MustGiveNonce(addr)
	if err != nil {
		return nil, err
	}
	tr := newTxRunner(nonce)
	txJob(tr)
	if tr.tx == nil && tr.err == nil {
		n.Impl.CommitNonce(addr, nonce, false)
		return nil, errors.New("must commit")
	}
	if err := tr.err; err != nil {
		n.Impl.CommitNonce(addr, nonce, false)
		if strings.Contains(err.Error(), "nonce too low") || strings.Contains(err.Error(), "nonce too high") {
			new_nonce, _ := n.Impl.SyncNonce(addr)
			log.Debugf("nonce:%d of %s is [%v], auto sync to %d", nonce, addr.Hex(), err, new_nonce)
		}
		return nil, err
	} else {
		n.Impl.CommitNonce(addr, nonce, true)
		return tr.tx, nil
	}
}
