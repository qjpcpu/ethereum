package contracts

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	//"math/big"
	"regexp"
	"strings"
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

func (tx *TransactionWithExtra) IsSuccess(conn *ethclient.Client) (bool, error) {
	rep, err := conn.TransactionReceipt(context.Background(), tx.Hash())
	if err != nil {
		return false, err
	}
	return rep.Status == types.ReceiptStatusSuccessful, nil
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

func DeployContract(conn *ethclient.Client, keyJson, keyPasswd, tokenABI, tokenBin string, params ...interface{}) (common.Address, *types.Transaction, error) {
	if !strings.HasPrefix(tokenBin, `0x`) {
		tokenBin = `0x` + tokenBin
	}
	parsed, err := abi.JSON(strings.NewReader(tokenABI))
	if err != nil {
		return common.Address{}, nil, err
	}
	address, tx, _, err := bind.DeployContract(BuildTransactOpts(keyJson, keyPasswd), parsed, common.FromHex(tokenBin), conn, params...)
	if err != nil {
		return common.Address{}, nil, err
	}
	return address, tx, nil
}

func BuildTransactOpts(keyJson, keyPasswd string) *bind.TransactOpts {
	addr := struct {
		Address string `json:"address"`
	}{}
	if err := json.Unmarshal([]byte(keyJson), &addr); err != nil {
		panic(fmt.Sprintf("build transactOpts fail:%v", err))
	}
	return &bind.TransactOpts{
		From:  common.HexToAddress(`0x` + addr.Address),
		Nonce: nil,
		Signer: func(signer types.Signer, addresses common.Address,
			transaction *types.Transaction) (*types.Transaction, error) {
			key, err := keystore.DecryptKey([]byte(keyJson), keyPasswd)
			if err != nil {
				return nil, err
			}
			signTransaction, err := types.SignTx(transaction, signer, key.PrivateKey)
			if err != nil {
				return nil, err
			}
			return signTransaction, nil
		},
		Value:   nil,
		Context: context.Background(),
	}
}

func Keccak256Hash(data string) common.Hash {
	return crypto.Keccak256Hash([]byte(data))
}
