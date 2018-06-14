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

func TestVerifySign(t *testing.T) {
	from := "0xE35f3e2A93322b61e5D8931f806Ff38F4a4F4D88"
	signHex := "0xbd5280c97dd1f069875a512364ca8470c0d928c4701b2a7a89c775478caf7ac670a908823fc75e65e400592d57e9649b2f1f57b4fd3980f02fe9a683e271a41c1b"
	msg := []byte("Hello world")
	if err := VerifySign(from, signHex, msg); err != nil {
		t.Fatal(err)
	}
}
