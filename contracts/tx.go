package contracts

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
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

// nonce and gasPrice are optional
func ResendTransaction(conn *ethclient.Client, tx *types.Transaction, signerFunc bind.SignerFn, nonce uint64, gasPrice *big.Int) (*types.Transaction, error) {
	from := NewTxExtra(tx).From()
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
	rawTx := types.NewTransaction(nonce, *tx.To(), tx.Value(), tx.Gas(), gasPrice, tx.Data())
	signedTx, err := signerFunc(types.HomesteadSigner{}, from, rawTx)
	if err != nil {
		return nil, err
	}
	if err = conn.SendTransaction(context.Background(), signedTx); err != nil {
		return nil, err
	}
	return signedTx, err
}

func TransferETH(conn *ethclient.Client, from, to common.Address, amount *big.Int, signerFunc bind.SignerFn, nonce uint64, gasPrice *big.Int, notes ...string) (*types.Transaction, error) {
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
	var gasLimit uint64 = 21000
	var data []byte
	if len(notes) > 0 && notes[0] != "" {
		data = packString(notes[0])
		msg := ethereum.CallMsg{From: from, To: &to, Value: amount, Data: data}
		var err error
		gasLimit, err = conn.EstimateGas(context.TODO(), msg)
		if err != nil {
			return nil, err
		}
	}
	rawTx := types.NewTransaction(nonce, to, amount, gasLimit, gasPrice, data)
	signedTx, err := signerFunc(types.HomesteadSigner{}, from, rawTx)
	if err != nil {
		return nil, err
	}
	if err = conn.SendTransaction(context.Background(), signedTx); err != nil {
		return nil, err
	}
	return signedTx, err
}

func packString(str string) []byte {
	l := len(str)
	return append(
		math.PaddedBigBytes(math.U256(big.NewInt(int64(l))), 32),
		common.RightPadBytes([]byte(str), (l+31)/32*32)...,
	)
}
