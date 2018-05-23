package unionpay

import (
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/qjpcpu/ethereum/contracts"
	"github.com/qjpcpu/ethereum/key"
	"math/big"
)

func PackPayParams(from common.Address, to common.Address, amount *big.Int, cutPercentage int, receiptId *big.Int, extra *big.Int) ([]byte, error) {
	if extra == nil {
		extra = big.NewInt(0)
	}
	if receiptId == nil {
		return nil, errors.New("no receipt id")
	}
	if cutPercentage < 0 || cutPercentage > 100 {
		return nil, errors.New("cutPercentage should in[0,100]")
	}
	msg := crypto.Keccak256(
		from.Bytes(),
		to.Bytes(),
		contracts.PackNum(amount),
		contracts.PackNum(big.NewInt(int64(cutPercentage))),
		contracts.PackNum(receiptId),
		contracts.PackNum(extra),
	)
	return msg, nil
}

func SignPayParams(keyjson, keypwd string, packedParams []byte) (string, error) {
	_, pk, err := key.ExportPrivateKey([]byte(keyjson), keypwd)
	if err != nil {
		return "", err
	}
	msg, err := key.Sign(pk, packedParams)
	if err != nil {
		return "", err
	}
	return hexutil.Encode(msg), nil
}

func PackAndSignPayParams(keyjson, keypwd string, from common.Address, to common.Address, amount *big.Int, cutPercentage int, receiptId *big.Int, extra *big.Int) (string, error) {
	data, err := PackPayParams(from, to, amount, cutPercentage, receiptId, extra)
	if err != nil {
		return "", err
	}
	return SignPayParams(keyjson, keypwd, data)
}
