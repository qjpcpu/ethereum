package contracts

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"strings"
	"time"
)

type TxOptsBuilder struct {
	opts *bind.TransactOpts
	Err  error
}

type TransactionWithExtra struct {
	*types.Transaction
}

func NewTxExtra(tx *types.Transaction) *TransactionWithExtra {
	return &TransactionWithExtra{Transaction: tx}
}

func (tx *TransactionWithExtra) From() common.Address {
	getSigner := func(trans *types.Transaction) types.Signer {
		v, _, _ := trans.RawSignatureValues()
		var isProtectedV bool
		for loop := true; loop; loop = false {
			if v.BitLen() <= 8 {
				vv := v.Uint64()
				isProtectedV = vv != 27 && vv != 28
				break
			}
			isProtectedV = true
		}
		if v.Sign() != 0 && isProtectedV {
			var chainId *big.Int
			for loop := true; loop; loop = false {
				if v.BitLen() <= 64 {
					vv := v.Uint64()
					if vv == 27 || vv == 28 {
						chainId = new(big.Int)
						break
					}
					chainId = new(big.Int).SetUint64((vv - 35) / 2)
					break
				}
				nv := new(big.Int).Sub(v, big.NewInt(35))
				chainId = nv.Div(nv, big.NewInt(2))
			}
			return types.NewEIP155Signer(chainId)
		} else {
			return types.HomesteadSigner{}
		}
	}
	signer := getSigner(tx.Transaction)
	from, err := types.Sender(signer, tx.Transaction)
	if err != nil {
		return common.Address{}
	}
	return from
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
	address := crypto.CreateAddress(tx.From(), tx.Nonce())
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

func Keccak256Hash(data string) common.Hash {
	return crypto.Keccak256Hash([]byte(data))
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
	opts := &bind.TransactOpts{
		From:  DecodeKeystoreAddress([]byte(keyJson)),
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
		Value:   big.NewInt(0),
		Context: context.Background(),
	}
	return opts
}

func DecodeKeystoreAddress(keyJsonStr []byte) common.Address {
	addr := struct {
		Address string `json:"address"`
	}{}
	if err := json.Unmarshal(keyJsonStr, &addr); err != nil {
		panic(fmt.Sprintf("parse address fail:%v", err))
	}
	if !strings.HasPrefix(addr.Address, "0x") {
		addr.Address = `0x` + addr.Address
	}
	return common.HexToAddress(addr.Address)
}

func NewTxOptsBuilder(keyJson, keyPwd string) *TxOptsBuilder {
	return &TxOptsBuilder{opts: BuildTransactOpts(keyJson, keyPwd)}
}

func (b *TxOptsBuilder) Get() *bind.TransactOpts {
	t := new(bind.TransactOpts)
	*t = *b.opts
	if b.opts.Nonce != nil {
		t.Nonce = new(big.Int).Set(b.opts.Nonce)
	}
	if b.opts.Value != nil {
		t.Value = new(big.Int).Set(b.opts.Value)
	}
	if b.opts.GasPrice != nil {
		t.GasPrice = new(big.Int).Set(b.opts.GasPrice)
	}
	return t
}

func (b *TxOptsBuilder) BuildValue(val *big.Int) *TxOptsBuilder {
	b.opts.Value = val
	return b
}

func (b *TxOptsBuilder) BuildNonce(nonce *big.Int) *TxOptsBuilder {
	b.opts.Nonce = nonce
	return b
}

func (b *TxOptsBuilder) PeekFrom() common.Address {
	return b.opts.From
}

func (b *TxOptsBuilder) BuildSuggestGasPrice(conn *ethclient.Client) *TxOptsBuilder {
	var err error
	b.opts.GasPrice, err = conn.SuggestGasPrice(context.TODO())
	if err != nil {
		b.Err = err
	}
	return b
}

func (b *TxOptsBuilder) BuildGasLimitMannual(limit uint64) *TxOptsBuilder {
	b.opts.GasLimit = limit
	return b
}

func (b *TxOptsBuilder) BuildGasPriceMannual(limit uint64) *TxOptsBuilder {
	b.opts.GasPrice = new(big.Int).SetUint64(limit)
	return b
}

// method是真实函数名称如erc20的transfer
func (b *TxOptsBuilder) BuildGasLimit(conn *ethclient.Client, contract_addr common.Address, abi_str string, method string, params ...interface{}) *TxOptsBuilder {
	parsed, err := abi.JSON(strings.NewReader(abi_str))
	if err != nil {
		b.Err = err
		return b
	}
	input, err := parsed.Pack(method, params...)
	msg := ethereum.CallMsg{From: b.opts.From, To: &contract_addr, Value: b.opts.Value, Data: input}
	limit, err := conn.EstimateGas(context.TODO(), msg)
	if err != nil {
		b.Err = err
	} else {
		b.opts.GasLimit = limit
	}
	return b
}

func WaitTxDone(conn *ethclient.Client, txhash common.Hash, timeout ...time.Duration) error {
	var timeoutDur time.Duration
	if len(timeout) > 0 {
		timeoutDur = timeout[0]
	} else {
		timeoutDur = time.Minute * 30
	}
	sleepSec := 3 * time.Second
	if sleepSec > timeoutDur {
		sleepSec = timeoutDur
	}

	timeoutCh := time.After(timeoutDur)
	for {
		rep, err := conn.TransactionReceipt(context.Background(), txhash)
		if err != nil {
			if err != ethereum.NotFound {
				return err
			}
		} else {
			if rep.Status == types.ReceiptStatusSuccessful {
				return nil
			} else {
				return errors.New("tx failed")
			}
		}
		select {
		case <-timeoutCh:
			return errors.New("wait timeout")
		case <-time.After(sleepSec):
		}
	}
}
