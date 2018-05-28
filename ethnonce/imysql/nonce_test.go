package imysql

import (
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/qjpcpu/ethereum/ethnonce"
	"testing"
)

func _testinit() *ethnonce.NonceManager {
	conn, _ := ethclient.Dial("http://localhost:18545")

	creator := PrepareMysqlManager("root:root@tcp(10.0.2.2:3306)/funny?charset=utf8", "mmtk")
	return creator.SetEthClient(conn).Build()
}

func _testteardown(mgr *ethnonce.NonceManager) {
	m := mgr.Impl.(*mysqlManager)
	st, _ := m.db.Prepare("truncate mmtk")
	st.Exec()
	mgr.Close()
}

func TestGiveCommit(t *testing.T) {
	mgr := _testinit()
	defer _testteardown(mgr)
	addr := common.HexToAddress(`0xe35f3e2a93322b61e5d8931f806ff38f4a4f4d88`)
	mgr.SyncNonce(addr)
	nonce, err := mgr.MustGiveNonce(addr)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("nonce:", nonce)
	if err = mgr.CommitNonce(addr, nonce, true); err != nil {
		t.Fatal(err)
	}
	nonce1, err := mgr.MustGiveNonce(addr)
	if err != nil {
		t.Fatal(err)
	}
	if nonce1 != nonce+1 {
		t.Fatal("bad ", nonce1)
	}
	_, err = mgr.GiveNonce(addr)
	if err == nil {
		t.Fatal("should err")
	}
	if err = mgr.CommitNonce(addr, nonce1, false); err != nil {
		t.Fatal(err)
	}
}

func TestWrap(t *testing.T) {
	mgr := _testinit()
	defer _testteardown(mgr)
	addr := common.HexToAddress(`0xe35f3e2a93322b61e5d8931f806ff38f4a4f4d88`)
	_, err := mgr.MustGiveNonce(addr)
	if err == nil {
		t.Fatal("should initial first")
	}
	mgr.SyncNonce(addr)
	nonce1 := mgr.PeekNonce(addr)
	_, err = mgr.GiveNonceForTx(addr, func(nonce uint64) (*types.Transaction, error) {
		return nil, errors.New("err")
	})
	if err == nil {
		t.Fatal("should error")
	}
	nonce2 := mgr.PeekNonce(addr)
	if nonce1 != nonce2 {
		t.Fatal("nonce should keep unchanged")
	}
	mgr.GiveNonceForTx(addr, func(nonce uint64) (*types.Transaction, error) {
		return new(types.Transaction), nil
	})
	nonce2 = mgr.PeekNonce(addr)
	if nonce1+1 != nonce2 {
		t.Fatal("nonce should increased")
	}
}
