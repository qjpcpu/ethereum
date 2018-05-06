package contracts

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
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
