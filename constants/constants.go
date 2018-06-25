package constants

import (
	"math/big"
)

const (
	Unit_1gwei uint64 = 1000000000          // 1 gwei
	Unit_1eth  uint64 = 1000000000000000000 // 1 eth
)

func WeiToEth(num *big.Int) *big.Float {
	one_eth := big.NewFloat(float64(Unit_1eth))
	return new(big.Float).Quo(new(big.Float).SetInt(num), one_eth)
}

func EthToWei(float_eth *big.Float) *big.Int {
	one_eth := big.NewFloat(float64(Unit_1eth))
	res := new(big.Int)
	new(big.Float).Mul(float_eth, one_eth).Int(res)
	return res
}

func WeiToGwei(num *big.Int) *big.Int {
	return new(big.Int).Quo(new(big.Int).Set(num), new(big.Int).SetUint64(Unit_1gwei))
}

// return n_gwei*1000000000 wei
func NGwei(n_gwei uint64) *big.Int {
	return new(big.Int).Mul(new(big.Int).SetUint64(Unit_1gwei), new(big.Int).SetUint64(n_gwei))
}
