package ethnonce

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/syndtr/goleveldb/leveldb"
	"strconv"
	"strings"
	"sync"
	"time"
)

type lvldbManager struct {
	*nonceCommon
	db *leveldb.DB
	*sync.Mutex
}

func newLvldbManager(nc *nonceCommon, file_path string) *lvldbManager {
	db, err := leveldb.OpenFile(file_path, nil)
	if err != nil {
		panic(err)
	}
	return &lvldbManager{
		nonceCommon: nc,
		db:          db,
		Mutex:       new(sync.Mutex),
	}
}

func (n *lvldbManager) PeekNonce(addr common.Address) uint64 {
	num, _ := resToNumber(n.db.Get([]byte(addressField(addr)), nil))
	return num
}

func (n *lvldbManager) GiveNonce(addr common.Address) (uint64, error) {
	n.Lock()
	defer n.Unlock()
	nonce, err := resToNumber(n.db.Get([]byte(addressField(addr)), nil))
	if err != nil {
		if err == leveldb.ErrNotFound {
			return 0, ErrNotInitAddress
		}
		return 0, ErrOtherHoldNonce
	}
	cmt, _ := resToNumber(n.db.Get([]byte(addressCommitField(addr)), nil))
	if cmt == 0 {
		batch := new(leveldb.Batch)
		batch.Put([]byte(addressCommitField(addr)), numberToString(1))
		batch.Put([]byte(addressTimestampField(addr)), numberToString(uint64(time.Now().Unix())))
		n.db.Write(batch, nil)
		return nonce, nil
	}
	last, _ := resToNumber(n.db.Get([]byte(addressTimestampField(addr)), nil))
	if time.Unix(int64(last), 0).Add(time.Second * 60).After(time.Now()) {
		return 0, ErrOtherHoldNonce
	}
	n.db.Put([]byte(addressTimestampField(addr)), numberToString(uint64(time.Now().Unix())), nil)
	return nonce, nil
}

func (n *lvldbManager) SyncNonce(addr common.Address) (uint64, error) {
	n.Lock()
	defer n.Unlock()
	cmt, err1 := resToNumber(n.db.Get([]byte(addressCommitField(addr)), nil))
	last, err2 := resToNumber(n.db.Get([]byte(addressTimestampField(addr)), nil))
	if err1 == nil && err2 == nil && cmt != 0 && time.Unix(int64(last), 0).Add(time.Second*60).After(time.Now()) {
		return 0, ErrOtherHoldNonce
	}
	nonce, err := n.ethConn.PendingNonceAt(context.Background(), addr)
	if err != nil {
		return 0, err
	}

	batch := new(leveldb.Batch)
	batch.Put([]byte(addressField(addr)), numberToString(nonce))
	batch.Put([]byte(addressCommitField(addr)), numberToString(0))
	batch.Put([]byte(addressTimestampField(addr)), numberToString(0))
	n.db.Write(batch, nil)

	return nonce, err
}

func (n *lvldbManager) CommitNonce(addr common.Address, nonce_number uint64, success bool) error {
	n.Lock()
	defer n.Unlock()
	dbnonce, err := resToNumber(n.db.Get([]byte(addressField(addr)), nil))
	if err != nil {
		if err == leveldb.ErrNotFound {
			return ErrNotInitAddress
		}
		return err
	}
	cmt, _ := resToNumber(n.db.Get([]byte(addressCommitField(addr)), nil))
	if cmt == 0 {
		return nil
	}
	if dbnonce != nonce_number {
		return ErrOtherHoldNonce
	}
	batch := new(leveldb.Batch)
	if success {
		batch.Put([]byte(addressField(addr)), numberToString(nonce_number+1))
	}
	batch.Put([]byte(addressCommitField(addr)), numberToString(0))
	batch.Put([]byte(addressTimestampField(addr)), numberToString(uint64(time.Now().Unix())))
	return n.db.Write(batch, nil)
}

func (n *lvldbManager) Close() error {
	return n.db.Close()
}

func resToNumber(data []byte, err error) (uint64, error) {
	if err != nil {
		return 0, err
	}
	return strconv.ParseUint(string(data), 10, 64)
}

func numberToString(num uint64) []byte {
	return []byte(fmt.Sprint(num))
}

func addressField(addr common.Address) string {
	return strings.ToLower(addr.Hex())
}

func addressCommitField(addr common.Address) string {
	return addressField(addr) + "_cmt"
}

func addressTimestampField(addr common.Address) string {
	return addressField(addr) + "_stp"
}
