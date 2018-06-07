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
	pk, _ := key.PrivateKeyFromHex("18641de16d5ef5fe9575769422e1138453ead48787fe54094d347483a5fa52b0")
	pwd := "123"
	platform, keyjson, _ := key.ImportPrivateKey(pk, pwd, keystore.LightScryptN, keystore.LightScryptP)
	from := common.HexToAddress("0xca35b7d915458ef540ade6068dfe2f44e8fa733c")
	to := common.HexToAddress("0xdd870fa1b7c4700f2bd7f44238821c26f7392148")
	amount := big.NewInt(1000000000000000000)
	cut := big.NewInt(45)
	nonce := big.NewInt(3)
	extra := big.NewInt(1)
	sign, err := PackAndSignPayParams(string(keyjson), pwd, from, to, amount, cut, nonce, extra)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("platform:", platform.Hex())
	t.Log("from:", from.Hex())
	t.Log("value:", amount.Uint64())
	t.Logf("safePay(\"%v\",%v,%v,%v,\"%v\")", to.Hex(), nonce.Uint64(), cut, extra.Uint64(), hexutil.Encode(sign))
	data, err := MakeSafePayTxData(string(keyjson), pwd, from, to, amount, int(cut.Int64()), nonce, extra)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("raw tx data:", data)
	data, err = MakeFixedSafePayTxData(string(keyjson), pwd, from, to, amount, cut, nonce, extra)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("raw tx data:", data)
}
