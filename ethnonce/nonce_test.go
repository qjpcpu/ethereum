package ethnonce

import (
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/garyburd/redigo/redis"
	"testing"
	"time"
)

func _testinit() *NonceManager {
	conn, _ := ethclient.Dial("http://localhost:18545")
	c := &redis.Pool{
		MaxIdle:     200,
		MaxActive:   200,
		IdleTimeout: 2 * time.Second,
		Dial: func() (redis.Conn, error) {
			connect_timeout := 2 * time.Second
			read_timeout := 2 * time.Second
			write_timeout := 2 * time.Second
			c, err := redis.DialTimeout("tcp", ":6379", connect_timeout,
				read_timeout, write_timeout)
			if err != nil {
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
	rc := c.Get()
	rc.Do("DEL", "testhash")
	rc.Close()
	return NewNonceManager(conn, c, "testhash")
}

func TestGiveCommit(t *testing.T) {
	mgr := _testinit()
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
