package contracts

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"regexp"
)

var regFrom = regexp.MustCompile(`From:\s+([^\s]+)`)

type TransactionWithExtra struct {
	*types.Transaction
}

func (tx *TransactionWithExtra) From() *common.Address {
	arr := regFrom.FindStringSubmatch(tx.Transaction.String())
	if len(arr) == 2 {
		addr := common.HexToAddress(arr[1])
		return &addr
	}
	return nil
}

// 某个地址是否合约
func IsContract(conn *ethclient.Client, hexAddr string) bool {
	code, err := conn.CodeAt(context.Background(), common.HexToAddress(hexAddr), nil)
	return err == nil && len(code) > 0
}
