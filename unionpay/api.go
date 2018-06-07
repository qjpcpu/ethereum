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

func PackPayParams(from common.Address, to common.Address, amount *big.Int, cut *big.Int, receiptId *big.Int, extra *big.Int) ([]byte, error) {
	if extra == nil {
		return nil, errors.New("no extra")
	}
	if receiptId == nil {
		return nil, errors.New("no receipt id")
	}
	if cut == nil {
		return nil, errors.New("no cut")
	}
	msg := crypto.Keccak256(
		from.Bytes(),
		to.Bytes(),
		contracts.PackNum(amount),
		contracts.PackNum(cut),
		contracts.PackNum(receiptId),
		contracts.PackNum(extra),
	)
	return msg, nil
}

func SignPayParams(keyjson, keypwd string, packedParams []byte) ([]byte, error) {
	_, pk, err := key.ExportPrivateKey([]byte(keyjson), keypwd)
	if err != nil {
		return nil, err
	}
	msg, err := key.Sign(pk, packedParams)
	if err != nil {
		return nil, err
	}
	return msg, nil
}

func PackAndSignPayParams(keyjson, keypwd string, from common.Address, to common.Address, amount *big.Int, cut *big.Int, receiptId *big.Int, extra *big.Int) ([]byte, error) {
	data, err := PackPayParams(from, to, amount, cut, receiptId, extra)
	if err != nil {
		return nil, err
	}
	return SignPayParams(keyjson, keypwd, data)
}

func makeUnionPayTxData(
	platform_keyjson,
	platform_keypwd string,
	from common.Address,
	to common.Address,
	amount *big.Int,
	cut *big.Int,
	receiptId *big.Int,
	extra *big.Int,
	method string,
) ([]byte, error) {
	sign, err := PackAndSignPayParams(platform_keyjson, platform_keypwd, from, to, amount, cut, receiptId, extra)
	if err != nil {
		return nil, err
	}
	_abi, err := contracts.ParseABI(UnionPayABI)
	if err != nil {
		return nil, err
	}
	return contracts.PackArguments(_abi, method, to, receiptId, cut, extra, sign)
}

func MakeSafePayTxData(
	platform_keyjson,
	platform_keypwd string,
	from common.Address,
	to common.Address,
	amount *big.Int,
	cutPercent int,
	receiptId *big.Int,
	extra *big.Int,
) (string, error) {
	cut := big.NewInt(int64(cutPercent))
	data, err := makeUnionPayTxData(platform_keyjson, platform_keypwd, from, to, amount, cut, receiptId, extra, "safePay")
	if err != nil {
		return "", err
	}
	return hexutil.Encode(data), nil
}

func MakeFixedSafePayTxData(
	platform_keyjson,
	platform_keypwd string,
	from common.Address,
	to common.Address,
	amount *big.Int,
	cutFixed *big.Int,
	receiptId *big.Int,
	extra *big.Int,
) (string, error) {
	data, err := makeUnionPayTxData(platform_keyjson, platform_keypwd, from, to, amount, cutFixed, receiptId, extra, "fixedSafePay")
	if err != nil {
		return "", err
	}
	return hexutil.Encode(data), nil
}
