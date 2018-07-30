package contracts

import (
	"context"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"time"
)

func MorningBlockOfDay(conn *ethclient.Client, blockno uint64) (uint64, error) {
	tblock, err := conn.HeaderByNumber(context.Background(), new(big.Int).SetUint64(blockno))
	if err != nil {
		return 0, err
	}
	bn, _, err := GetBlockInTimeMinute(conn, time.Unix(tblock.Time.Int64(), 0).Truncate(24*time.Hour), tblock)
	if err != nil {
		return 0, err
	}
	return bn, nil
}

func GetBlockInTimeMinute(conn *ethclient.Client, targetTime time.Time, baseBlock *types.Header) (uint64, time.Time, error) {
	const secondsPerBlock uint64 = 15
	getBlockTime := func(bn uint64) (time.Time, error) {
		tblock, err := conn.HeaderByNumber(context.Background(), new(big.Int).SetUint64(bn))
		if err != nil {
			return time.Time{}, err
		}
		return time.Unix(tblock.Time.Int64(), 0), nil
	}
	var tm time.Time
	var err error
	var blockno uint64
	if baseBlock != nil {
		tm = time.Unix(baseBlock.Time.Int64(), 0)
		blockno = baseBlock.Number.Uint64()
	} else {
		b, err := conn.HeaderByNumber(context.Background(), nil)
		if err != nil {
			return 0, time.Time{}, err
		}
		tm = time.Unix(b.Time.Int64(), 0)
		blockno = b.Number.Uint64()
	}
	morning := targetTime.Truncate(60 * time.Second)
	for {
		if tm.After(morning) && tm.Before(morning.Add(60*time.Second)) {
			break
		}
		var offset uint64
		if tm.After(morning) {
			offset = uint64(tm.Sub(morning).Seconds()) / secondsPerBlock
			if offset == 0 {
				blockno -= 1
			} else {
				blockno -= offset
			}
		} else {
			offset = (uint64(morning.Sub(tm).Seconds()) + secondsPerBlock) / secondsPerBlock
			if offset == 0 {
				blockno += 1
			} else {
				blockno += offset
			}
		}
		if tm, err = getBlockTime(blockno); err != nil {
			return 0, time.Time{}, err
		}
	}
	return blockno, tm, nil
}
