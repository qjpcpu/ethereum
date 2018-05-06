package ethnonce

import (
	"context"
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/garyburd/redigo/redis"
	"strings"
	"time"
)

var (
	ErrNotInitAddress = errors.New("not initailize")
	ErrOtherHoldNonce = errors.New("others hold the nonce")

	// redis-cli --eval give.lua hash_name , 0x123 timestamp(in second)
	giveNonceScript = redis.NewScript(1, `
local key = KEYS[1]
local field_addr = ARGV[1]
local timestamp = tonumber(ARGV[2])  -- second
local field_commit = field_addr.."_cmt"
local field_timestamp = field_addr.."_stp"
-- addr not exist
if redis.call("HEXISTS", key, field_addr) == 0 then
	return -1
end
local iscommit = tonumber(redis.call("HGET", key, field_commit))
-- 0: can give nonce
if iscommit == 0 then
	redis.call("HSET", key, field_timestamp, timestamp)
	redis.call("HSET", key, field_commit, 1)
	return redis.call("HGET", key, field_addr)
end
-- come here,means user don't commit
-- user not commit, but 1 minute timeout, so we still give nonce
if tonumber(redis.call("HGET",key, field_timestamp)) + 60 < timestamp then
	redis.call("HSET", key, field_timestamp, timestamp)
	return redis.call("HGET", key, field_addr)
end

-- user hold nonce, we should wait
return -2
`)
	// redis-cli --eval ./comit_nonce.lua n , 0x123 timestamp nonce 0/1(成功使用/放弃)
	commitNonceScript = redis.NewScript(1, `
local key = KEYS[1]
local field_addr = ARGV[1]
local timestamp = tonumber(ARGV[2])  -- second
local nonce = tonumber(ARGV[3])
local commit_ok = tonumber(ARGV[4])  -- 0/1 0: ok
local field_commit = field_addr.."_cmt"
local field_timestamp = field_addr.."_stp"
-- addr not exist
if redis.call("HEXISTS", key, field_addr) == 0 then
	return -1
end
local iscommit = tonumber(redis.call("HGET", key, field_commit))
-- iscommit==1 means can commit
if iscommit ~= 0 then
	redis.call("HSET", key, field_timestamp, timestamp)
	redis.call("HSET", key, field_commit, 0)
    local n = tonumber(redis.call("HGET",key, field_addr))
    if n ~= nonce then
        return -2
    end
	if commit_ok == 0 then
	    redis.call("HINCRBY",key, field_addr, 1)
    end
	return 0
end

-- iscommit already == 0, we consider is OK
return 0
`)
	// redis-cli --eval ./sync_nonce.lua n , 0x123 nonce_number timestamp
	syncNonceScript = redis.NewScript(1, `
local key = KEYS[1]
local field_addr = ARGV[1]
local nonce = ARGV[2]
local timestamp = tonumber(ARGV[3])  -- second
local field_commit = field_addr.."_cmt"
local field_timestamp = field_addr.."_stp"

if redis.call("HEXISTS", key, field_addr) ~= 0 then
    local iscommit = tonumber(redis.call("HGET", key, field_commit))
    local lasttime = tonumber(redis.call("HGET",key, field_timestamp))
    if iscommit ~= 0 and lasttime + 60 > timestamp then
        return -1
    end
end
redis.call("HSET", key, field_addr, nonce)
redis.call("HSET", key, field_timestamp, 0)
redis.call("HSET", key, field_commit, 0)
return 0
`)
)

type NonceManager struct {
	NoncesHash string
}

func NewNonceManager(nonce_name string) NonceManager {
	return NonceManager{NoncesHash: nonce_name}
}

func (n NonceManager) MustGiveNonce(conn redis.Conn, addr common.Address, ethConn ...*ethclient.Client) (uint64, error) {
	var code uint64
	var err error
	for i := 0; i < 600; i++ {
		code, err = n.GiveNonce(conn, addr, ethConn...)
		if err == nil || err == ErrNotInitAddress {
			break
		}
		// err == ErrOtherHoldNonce
		<-time.After(100 * time.Millisecond)
	}
	return code, err
}

func (n NonceManager) GiveNonce(conn redis.Conn, addr common.Address, ethConn ...*ethclient.Client) (uint64, error) {
	address := strings.ToLower(addr.Hex())
	now := time.Now()
	errcode, err := redis.Int64(giveNonceScript.Do(conn, n.NoncesHash, address, now.Unix()))
	if err != nil && err != redis.ErrNil {
		return 0, err
	}
	if len(ethConn) > 0 && ethConn[0] != nil && errcode == -1 {
		return n.SyncNonce(conn, addr, ethConn[0])
	}
	switch errcode {
	case -1:
		return 0, ErrNotInitAddress
	case -2:
		return 0, ErrOtherHoldNonce
	default:
		return uint64(errcode), nil
	}
}

func (n NonceManager) SyncNonce(redis_conn redis.Conn, addr common.Address, conn *ethclient.Client) (uint64, error) {
	nonce, err := conn.PendingNonceAt(context.Background(), addr)
	if err != nil {
		return 0, err
	}
	code, err := redis.Int(syncNonceScript.Do(redis_conn, n.NoncesHash, strings.ToLower(addr.Hex()), nonce, time.Now().Unix()))
	if err != nil {
		return 0, err
	}
	if code == -1 {
		return 0, ErrOtherHoldNonce
	}
	return nonce, err
}

func (n NonceManager) CommitNonce(redis_conn redis.Conn, addr common.Address, nonce_number uint64, success bool) error {
	ok := 0
	if !success {
		ok = 1
	}
	code, err := redis.Int(commitNonceScript.Do(redis_conn, n.NoncesHash, strings.ToLower(addr.Hex()), time.Now().Unix(), nonce_number, ok))
	if err != nil {
		return err
	}
	switch code {
	case -1:
		return ErrNotInitAddress
	case -2:
		return ErrOtherHoldNonce
	default:
		return nil
	}
}

func (n NonceManager) GiveNonceForTx(eth_conn *ethclient.Client, redis_conn redis.Conn, addr common.Address, txJob func(nonce uint64) (*types.Transaction, error)) (*types.Transaction, error) {
	nonce, err := n.MustGiveNonce(redis_conn, addr)
	if err != nil {
		return nil, err
	}
	if tx, err := txJob(nonce); err != nil {
		n.CommitNonce(redis_conn, addr, nonce, false)
		if err == core.ErrNonceTooLow {
			n.SyncNonce(redis_conn, addr, eth_conn)
		}
		return nil, err
	} else {
		n.CommitNonce(redis_conn, addr, nonce, true)
		return tx, nil
	}
}
