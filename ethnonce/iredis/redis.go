package iredis

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/garyburd/redigo/redis"
	"github.com/qjpcpu/ethereum/ethnonce"
	"strings"
	"time"
)

var (
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

type redisManager struct {
	NoncesName string
	ethConn    *ethclient.Client
	pool       *redis.Pool
}

type RedisManagerCreator struct {
	mgr *redisManager
}

func (rc *RedisManagerCreator) SetEthClient(conn *ethclient.Client) ethnonce.ManagerCreator {
	rc.mgr.ethConn = conn
	return rc
}

func (rc *RedisManagerCreator) Build() *ethnonce.NonceManager {
	return &ethnonce.NonceManager{
		Impl: rc.mgr,
	}
}

func PrepareRedisManager(key string, conn string, redis_db, passwd string) ethnonce.ManagerCreator {
	pool := &redis.Pool{
		MaxIdle:     200,
		MaxActive:   200,
		IdleTimeout: 2 * time.Second,
		Dial: func() (redis.Conn, error) {
			connect_timeout := 2 * time.Second
			read_timeout := 2 * time.Second
			write_timeout := 2 * time.Second
			c, err := redis.DialTimeout("tcp", conn, connect_timeout,
				read_timeout, write_timeout)
			if err != nil {
				return nil, err
			}

			if passwd != "" {
				if _, err := c.Do("AUTH", passwd); err != nil {
					c.Close()
					return nil, err
				}
			}

			if redis_db != "" {
				if _, err = c.Do("SELECT", redis_db); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
	return PrepareRedisPoolManager(key, pool)
}

func PrepareRedisPoolManager(key string, pool *redis.Pool) ethnonce.ManagerCreator {
	return &RedisManagerCreator{
		mgr: &redisManager{
			NoncesName: key,
			pool:       pool,
		},
	}
}

func (n *redisManager) PeekNonce(addr common.Address) uint64 {
	conn := n.pool.Get()
	defer conn.Close()
	num, _ := redis.Uint64(conn.Do("HGET", n.NoncesName, strings.ToLower(addr.Hex())))
	return num
}

func (n *redisManager) GiveNonce(addr common.Address) (uint64, error) {
	conn := n.pool.Get()
	defer conn.Close()
	address := strings.ToLower(addr.Hex())
	now := time.Now()
	errcode, err := redis.Int64(giveNonceScript.Do(conn, n.NoncesName, address, now.Unix()))
	if err != nil && err != redis.ErrNil {
		return 0, err
	}
	switch errcode {
	case -1:
		return 0, ethnonce.ErrNotInitAddress
	case -2:
		return 0, ethnonce.ErrOtherHoldNonce
	default:
		return uint64(errcode), nil
	}
}

func (n *redisManager) SyncNonce(addr common.Address) (uint64, error) {
	redis_conn := n.pool.Get()
	defer redis_conn.Close()
	nonce, err := n.ethConn.PendingNonceAt(context.Background(), addr)
	if err != nil {
		return 0, err
	}
	code, err := redis.Int(syncNonceScript.Do(redis_conn, n.NoncesName, strings.ToLower(addr.Hex()), nonce, time.Now().Unix()))
	if err != nil {
		return 0, err
	}
	if code == -1 {
		return 0, ethnonce.ErrOtherHoldNonce
	}
	return nonce, err
}

func (n *redisManager) CommitNonce(addr common.Address, nonce_number uint64, success bool) error {
	redis_conn := n.pool.Get()
	defer redis_conn.Close()
	ok := 0
	if !success {
		ok = 1
	}
	code, err := redis.Int(commitNonceScript.Do(redis_conn, n.NoncesName, strings.ToLower(addr.Hex()), time.Now().Unix(), nonce_number, ok))
	if err != nil {
		return err
	}
	switch code {
	case -1:
		return ethnonce.ErrNotInitAddress
	case -2:
		return ethnonce.ErrOtherHoldNonce
	default:
		return nil
	}
}

func (n *redisManager) Close() error {
	return n.pool.Close()
}
