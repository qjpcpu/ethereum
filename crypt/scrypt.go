package crypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/randentropy"
	"golang.org/x/crypto/scrypt"
)

const (
	scryptR     = 8
	scryptDKLen = 32
)

func Encrypt(keyBytes []byte, auth string) ([]byte, error) {
	return EncryptNP(keyBytes, auth, keystore.LightScryptN, keystore.LightScryptP)
}

func EncryptStandard(keyBytes []byte, auth string) ([]byte, error) {
	return EncryptNP(keyBytes, auth, keystore.StandardScryptN, keystore.StandardScryptP)
}

func EncryptNP(keyBytes []byte, auth string, scryptN, scryptP int) ([]byte, error) {
	authArray := []byte(auth)
	salt := randentropy.GetEntropyCSPRNG(32)
	derivedKey, err := scrypt.Key(authArray, salt, scryptN, scryptR, scryptP, scryptDKLen)
	if err != nil {
		return nil, err
	}
	encryptKey := derivedKey[:16]

	iv := randentropy.GetEntropyCSPRNG(aes.BlockSize) // 16
	cipherText, err := aesCTRXOR(encryptKey, keyBytes, iv)
	if err != nil {
		return nil, err
	}
	mac := crypto.Keccak256(derivedKey[16:32], cipherText)

	scryptParamsJSON := make(map[string]interface{}, 5)
	scryptParamsJSON["n"] = scryptN
	scryptParamsJSON["r"] = scryptR
	scryptParamsJSON["p"] = scryptP
	scryptParamsJSON["dklen"] = scryptDKLen
	scryptParamsJSON["salt"] = hex.EncodeToString(salt)

	cipherParamsJSON := cipherparamsJSON{
		IV: hex.EncodeToString(iv),
	}

	keyHeaderKDF := "scrypt"
	cryptoStruct := cryptoJSON{
		Cipher:       "aes-128-ctr",
		CipherText:   hex.EncodeToString(cipherText),
		CipherParams: cipherParamsJSON,
		KDF:          keyHeaderKDF,
		KDFParams:    scryptParamsJSON,
		MAC:          hex.EncodeToString(mac),
	}
	return json.Marshal(cryptoStruct)
}

func aesCTRXOR(key, inText, iv []byte) ([]byte, error) {
	// AES-128 is selected due to size of encryptKey.
	aesBlock, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	stream := cipher.NewCTR(aesBlock, iv)
	outText := make([]byte, len(inText))
	stream.XORKeyStream(outText, inText)
	return outText, err
}

type cryptoJSON struct {
	Cipher       string                 `json:"cipher"`
	CipherText   string                 `json:"ciphertext"`
	CipherParams cipherparamsJSON       `json:"cipherparams"`
	KDF          string                 `json:"kdf"`
	KDFParams    map[string]interface{} `json:"kdfparams"`
	MAC          string                 `json:"mac"`
}

type cipherparamsJSON struct {
	IV string `json:"iv"`
}

func Decrypt(keyjson []byte, auth string) ([]byte, error) {
	k := new(cryptoJSON)
	if err := json.Unmarshal(keyjson, k); err != nil {
		return nil, err
	}
	return decryptKeyV3(k, auth)

}

func decryptKeyV3(crptojs *cryptoJSON, auth string) (keyBytes []byte, err error) {
	if crptojs.Cipher != "aes-128-ctr" {
		return nil, fmt.Errorf("Cipher not supported: %v", crptojs.Cipher)
	}

	mac, err := hex.DecodeString(crptojs.MAC)
	if err != nil {
		return nil, err
	}

	iv, err := hex.DecodeString(crptojs.CipherParams.IV)
	if err != nil {
		return nil, err
	}

	cipherText, err := hex.DecodeString(crptojs.CipherText)
	if err != nil {
		return nil, err
	}

	derivedKey, err := getKDFKey(*crptojs, auth)
	if err != nil {
		return nil, err
	}

	calculatedMAC := crypto.Keccak256(derivedKey[16:32], cipherText)
	if !bytes.Equal(calculatedMAC, mac) {
		return nil, errors.New("could not decrypt content with given passphrase")
	}

	plainText, err := aesCTRXOR(derivedKey[:16], cipherText, iv)
	if err != nil {
		return nil, err
	}
	return plainText, err
}

func getKDFKey(cryptoJSON cryptoJSON, auth string) ([]byte, error) {
	authArray := []byte(auth)
	salt, err := hex.DecodeString(cryptoJSON.KDFParams["salt"].(string))
	if err != nil {
		return nil, err
	}
	dkLen := ensureInt(cryptoJSON.KDFParams["dklen"])

	n := ensureInt(cryptoJSON.KDFParams["n"])
	r := ensureInt(cryptoJSON.KDFParams["r"])
	p := ensureInt(cryptoJSON.KDFParams["p"])
	return scrypt.Key(authArray, salt, n, r, p, dkLen)
}

func ensureInt(x interface{}) int {
	res, ok := x.(int)
	if !ok {
		res = int(x.(float64))
	}
	return res
}
