package key

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func VerifySign(from, sigHex string, msg []byte) error {
	fromAddr := common.HexToAddress(from)

	sig := hexutil.MustDecode(sigHex)
	if len(sig) != 65 {
		return fmt.Errorf("signature must be 65 bytes long")
	}
	// https://github.com/ethereum/go-ethereum/blob/55599ee95d4151a2502465e0afc7c47bd1acba77/internal/ethapi/api.go#L442
	if sig[64] != 27 && sig[64] != 28 {
		return fmt.Errorf("invalid Ethereum signature (V is not 27 or 28)")
	}
	sig[64] -= 27

	pubKey, err := crypto.SigToPub(signHash(msg), sig)
	if err != nil {
		return fmt.Errorf("sigToPub error:%v", err)
	}

	recoveredAddr := crypto.PubkeyToAddress(*pubKey)

	if fromAddr != recoveredAddr {
		return fmt.Errorf("recover address is not from address")
	}
	return nil
}

// https://github.com/ethereum/go-ethereum/blob/55599ee95d4151a2502465e0afc7c47bd1acba77/internal/ethapi/api.go#L404
// signHash is a helper function that calculates a hash for the given message that can be
// safely used to calculate a signature from.
//
// The hash is calculated as
//   keccak256("\x19Ethereum Signed Message:\n"${message length}${message}).
//
// This gives context to the signed message and prevents signing of transactions.
func signHash(data []byte) []byte {
	msg := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(data), data)
	return crypto.Keccak256([]byte(msg))
}

func Sign(privateKey *ecdsa.PrivateKey, msg []byte) ([]byte, error) {
	sig, err := crypto.Sign(signHash(msg), privateKey)
	if err != nil {
		return nil, err
	}
	sig[64] += 27 // Transform V from 0/1 to 27/28 according to the yellow paper
	return sig, nil
}
