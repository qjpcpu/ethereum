package key

import (
	crand "crypto/rand"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"testing"
)

func TestSignature(t *testing.T) {
	pk, err := newKey(crand.Reader)
	if err != nil {
		t.Fatal(err)
	}
	msg := "JasonGeek"
	sign, err := Sign(pk, []byte(msg))
	if err != nil {
		t.Fatal(err)
	}
	from := crypto.PubkeyToAddress(pk.PublicKey).Hex()
	signHex := hexutil.Encode(sign)
	if err := VerifySign(from, signHex, []byte(msg)); err != nil {
		t.Fatal(err)
	}
}
