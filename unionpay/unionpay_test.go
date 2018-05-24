package unionpay

import (
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/qjpcpu/ethereum/key"
	"math/big"
	"testing"
)

func TestSignature(t *testing.T) {
	pk, _ := key.StringToPrivateKey("18641de16d5ef5fe9575769422e1138453ead48787fe54094d347483a5fa52b0")
	pwd := "123"
	platform, keyjson, _ := key.ImportPrivateKey(pk, pwd, keystore.LightScryptN, keystore.LightScryptP)
	from := common.HexToAddress("0xca35b7d915458ef540ade6068dfe2f44e8fa733c")
	to := common.HexToAddress("0xdd870fa1b7c4700f2bd7f44238821c26f7392148")
	amount := big.NewInt(1000000000000000000)
	cut := 50
	nonce := big.NewInt(1)
	state := big.NewInt(522)
	sign, err := PackAndSignPayParams(string(keyjson), pwd, from, to, amount, cut, nonce, state)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("platform:", platform.Hex())
	t.Log("from:", from.Hex())
	t.Log("value:", amount.Uint64())
	t.Logf("payCash(%v,%v,\"%v\",%v,\"%v\")", nonce.Uint64(), cut, to.Hex(), state.Uint64(), hexutil.Encode(sign))
	data, err := MakeUnionPayTxData(string(keyjson), pwd, from, to, amount, cut, nonce, state)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("raw tx data:", hexutil.Encode(data))
}
