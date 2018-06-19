package contracts

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
)

type Address common.Address
type Hash common.Hash

func (a Address) MarshalJSON() ([]byte, error) {
	return json.Marshal(common.Address(a).Hex())
}

func (a *Address) UnmarshalJSON(b []byte) error {
	var str string
	json.Unmarshal(b, &str)
	*a = Address(common.HexToAddress(str))
	return nil
}

func (a Address) String() string {
	return a.Hex()
}

func (a Address) Hex() string {
	return a.StdAddress().Hex()
}

func (a Address) StdAddress() common.Address {
	return common.Address(a)
}

func (a Hash) MarshalJSON() ([]byte, error) {
	return json.Marshal(common.Hash(a).Hex())
}

func (a *Hash) UnmarshalJSON(b []byte) error {
	var str string
	json.Unmarshal(b, &str)
	*a = Hash(common.HexToHash(str))
	return nil
}

func (a Hash) String() string {
	return a.Hex()
}

func (a Hash) Hex() string {
	return a.StdHash().Hex()
}

func (a Hash) StdHash() common.Hash {
	return common.Hash(a)
}
