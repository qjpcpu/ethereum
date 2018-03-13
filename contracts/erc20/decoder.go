package erc20

import (
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
)

var (
	// a9059cbb
	transferFuncSig string = crypto.Keccak256Hash([]byte("transfer(address,uint256)")).Hex()[2:10]
	// 23b872dd
	transferFromFuncSig string = crypto.Keccak256Hash([]byte("transferFrom(address,address,uint256)")).Hex()[2:10]
)

var (
	NotTransferFuncErr     = errors.New("not transfer function")
	NotTransferFromFuncErr = errors.New("not transferFrom function")
	BadTransactionErr      = errors.New("bad transaction data")
)

func DecodeTransferData(txData []byte) (to common.Address, amount *big.Int, err error) {
	if len(txData) != 68 {
		err = BadTransactionErr
		return
	}
	mid, toAddr, amt := txData[:4], txData[4:36], txData[36:]
	if common.Bytes2Hex(mid) != transferFuncSig {
		err = NotTransferFuncErr
		return
	}
	to = common.BytesToAddress(toAddr)
	amount = new(big.Int).SetBytes(amt)
	return
}

func DecodeTransferFromData(txData []byte) (from common.Address, to common.Address, amount *big.Int, err error) {
	if len(txData) != 100 {
		err = BadTransactionErr
		return
	}
	mid, fromAddr, toAddr, amt := txData[:4], txData[4:36], txData[36:68], txData[68:]
	if common.Bytes2Hex(mid) != transferFromFuncSig {
		err = NotTransferFromFuncErr
		return
	}
	from = common.BytesToAddress(fromAddr)
	to = common.BytesToAddress(toAddr)
	amount = new(big.Int).SetBytes(amt)
	return
}

func IsTransferFunc(txData []byte) bool {
	return len(txData) > 4 && common.Bytes2Hex(txData[:4]) == transferFuncSig
}

func IsTransferFromFunc(txData []byte) bool {
	return len(txData) > 4 && common.Bytes2Hex(txData[:4]) == transferFromFuncSig
}
