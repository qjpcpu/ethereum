package contracts

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"strings"
)

func SignerFuncOf(keyJson, keyPasswd string) bind.SignerFn {
	return func(signer types.Signer, addresses common.Address,
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
	}
}

func SignerFuncOfPK(pk *ecdsa.PrivateKey) bind.SignerFn {
	return func(signer types.Signer, addresses common.Address, transaction *types.Transaction) (*types.Transaction, error) {
		signTransaction, err := types.SignTx(transaction, signer, pk)
		if err != nil {
			return nil, err
		}
		return signTransaction, nil
	}
}

// nonce and gasPrice are optional
func ResendTransaction(conn *ethclient.Client, tx *types.Transaction, signerFunc bind.SignerFn, nonce uint64, gasPrice *big.Int) (*types.Transaction, error) {
	return SendRawTransaction(
		conn,
		NewTxExtra(tx).From(),
		*tx.To(),
		tx.Value(),
		tx.Data(),
		signerFunc,
		nonce,
		gasPrice,
		tx.Gas(),
	)
}

// transfer eth with optinal note text
func TransferETH(conn *ethclient.Client, from, to common.Address, amount *big.Int, signerFunc bind.SignerFn, nonce uint64, gasPrice *big.Int, notes ...string) (*types.Transaction, error) {
	var gasLimit uint64 = 21000
	var data []byte
	if len(notes) > 0 && notes[0] != "" {
		data = PackString(notes[0])
		gasLimit = 0
	}
	return SendRawTransaction(conn, from, to, amount, data, signerFunc, nonce, gasPrice, gasLimit)
}

func PackString(str string) []byte {
	l := len(str)
	return append(
		math.PaddedBigBytes(math.U256(big.NewInt(int64(l))), 32),
		common.RightPadBytes([]byte(str), (l+31)/32*32)...,
	)
}

func PackNum(num *big.Int) []byte {
	return math.PaddedBigBytes(math.U256(new(big.Int).Set(num)), 32)
}

func PackAddress(addr common.Address) []byte {
	return common.LeftPadBytes(addr.Bytes(), 32)
}

func ParseABI(rawjson string) (abi.ABI, error) {
	return abi.JSON(strings.NewReader(rawjson))
}

func PackArguments(_abi abi.ABI, method string, params ...interface{}) ([]byte, error) {
	input, err := _abi.Pack(method, params...)
	if err != nil {
		return nil, err
	}
	return input, nil
}

func PackArgumentsWithNumber(_abi abi.ABI, method string, params ...interface{}) ([]byte, error) {
	method_params := params[:len(params)-1]
	num, ok := params[len(params)-1].(*big.Int)
	if !ok {
		return nil, errors.New("the last parameter should be *big.Int")
	}
	input, err := PackArguments(_abi, method, method_params...)
	if err != nil {
		return nil, err
	}
	input = append(input, PackNum(num)...)
	return input, nil
}

func PackArgumentsWithString(_abi abi.ABI, method string, params ...interface{}) ([]byte, error) {
	method_params := params[:len(params)-1]
	str, ok := params[len(params)-1].(string)
	if !ok {
		return nil, errors.New("the last parameter should be string")
	}
	input, err := PackArguments(_abi, method, method_params...)
	if err != nil {
		return nil, err
	}
	if str != "" {
		input = append(input, PackString(str)...)
	}
	return input, nil
}

func SendRawTransaction(conn *ethclient.Client, from, to common.Address, value *big.Int, data []byte, signerFunc bind.SignerFn, nonce uint64, gasPrice *big.Int, gasLimit uint64) (*types.Transaction, error) {
	if nonce == 0 {
		nonceInt, err := conn.PendingNonceAt(context.Background(), from)
		if err != nil {
			return nil, fmt.Errorf("failed to retrieve account nonce: %v", err)
		}
		nonce = nonceInt
	}
	if gasPrice == nil {
		gp, err := conn.SuggestGasPrice(context.Background())
		if err != nil {
			return nil, fmt.Errorf("failed to retrieve gas price: %v", err)
		}
		gasPrice = gp
	}
	if gasLimit == 0 {
		msg := ethereum.CallMsg{From: from, To: &to, Value: value, Data: data}
		var err error
		gasLimit, err = conn.EstimateGas(context.TODO(), msg)
		if err != nil {
			return nil, err
		}
	}
	rawTx := types.NewTransaction(nonce, to, value, gasLimit, gasPrice, data)
	signedTx, err := signerFunc(types.HomesteadSigner{}, from, rawTx)
	if err != nil {
		return nil, err
	}
	if err = conn.SendTransaction(context.Background(), signedTx); err != nil {
		return nil, err
	}
	return signedTx, err
}
