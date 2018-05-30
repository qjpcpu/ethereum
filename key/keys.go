package key

import (
	"crypto/ecdsa"
	crand "crypto/rand"
	"encoding/hex"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
	"io"
)

func NewPrivateKey() (*ecdsa.PrivateKey, error) {
	return newKey(crand.Reader)
}

func NewLightKey(passphrase string) (common.Address, []byte, error) {
	pk, err := newKey(crand.Reader)
	if err != nil {
		return common.Address{}, nil, err
	}
	return ImportPrivateKey(pk, passphrase, keystore.LightScryptN, keystore.LightScryptP)
}

func NewStandardKey(passphrase string) (common.Address, []byte, error) {
	pk, err := newKey(crand.Reader)
	if err != nil {
		return common.Address{}, nil, err
	}
	return ImportPrivateKey(pk, passphrase, keystore.StandardScryptN, keystore.StandardScryptP)
}

func ImportPrivateKeyLight(priv_key *ecdsa.PrivateKey, passphrase string) (common.Address, []byte, error) {
	return ImportPrivateKey(priv_key, passphrase, keystore.LightScryptN, keystore.LightScryptP)
}

func ImportPrivateKeyStandard(priv_key *ecdsa.PrivateKey, passphrase string) (common.Address, []byte, error) {
	return ImportPrivateKey(priv_key, passphrase, keystore.StandardScryptN, keystore.StandardScryptP)
}

func ImportPrivateKey(priv_key *ecdsa.PrivateKey, passphrase string, scryptN, scryptP int) (common.Address, []byte, error) {
	nkey := newKeyFromECDSA(priv_key)
	keyjson, err := keystore.EncryptKey(nkey, passphrase, scryptN, scryptP)
	if err != nil {
		return common.Address{}, nil, err
	}
	return crypto.PubkeyToAddress(nkey.PrivateKey.PublicKey), keyjson, nil
}

func ExportPrivateKey(keyjson []byte, auth string) (common.Address, *ecdsa.PrivateKey, error) {
	nkey, err := keystore.DecryptKey(keyjson, auth)
	if err != nil {
		return common.Address{}, nil, err
	}
	return nkey.Address, nkey.PrivateKey, nil
}

func PrivateKeyToHex(priv_key *ecdsa.PrivateKey) string {
	return hex.EncodeToString(crypto.FromECDSA(priv_key))
}

func PrivateKeyFromHex(str string) (*ecdsa.PrivateKey, error) {
	data, err := hex.DecodeString(str)
	if err != nil {
		return nil, err
	}
	priv, err := crypto.ToECDSA(data)
	if err != nil {
		return nil, err
	}
	return priv, nil
}

func newKey(rand io.Reader) (*ecdsa.PrivateKey, error) {
	return ecdsa.GenerateKey(crypto.S256(), rand)
}

func newKeyFromECDSA(privateKeyECDSA *ecdsa.PrivateKey) *keystore.Key {
	key := &keystore.Key{
		Address:    crypto.PubkeyToAddress(privateKeyECDSA.PublicKey),
		PrivateKey: privateKeyECDSA,
	}
	return key
}

func PrivateKeyToBytes(pk *ecdsa.PrivateKey) []byte {
	return math.PaddedBigBytes(pk.D, 32)
}

func PrivateKeyFromBytes(data []byte) *ecdsa.PrivateKey {
	return crypto.ToECDSAUnsafe(data)
}
