package contracts

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// 某个地址是否合约地址
func IsContract(conn *ethclient.Client, hexAddr string) bool {
	code, err := conn.CodeAt(context.Background(), common.HexToAddress(hexAddr), nil)
	return err == nil && len(code) > 0
}
