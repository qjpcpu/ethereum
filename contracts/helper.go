package contracts

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
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

// 对于合约创建交易,获取该交易创建的合约地址
func (tx *TransactionWithExtra) ContractAddress() *common.Address {
	address := crypto.CreateAddress(*tx.From(), tx.Nonce())
	return &address
}

// 某个地址是否合约
func IsContract(conn *ethclient.Client, hexAddr string) bool {
	code, err := conn.CodeAt(context.Background(), common.HexToAddress(hexAddr), nil)
	return err == nil && len(code) > 0
}

func (tx *TransactionWithExtra) IsContractCreation() bool {
	return tx.To() == nil
}
