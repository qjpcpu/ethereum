package contracts

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"testing"
)

func TestIsContract(t *testing.T) {
	addr := `0x86fa049857e0209aa7d9e616f7eb3b3b78ecfdb0` // EOS contract
	conn, err := ethclient.Dial("https://api.myetherapi.com/eth")
	if err != nil {
		t.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	isCon := IsContract(conn, addr)
	if !isCon {
		t.Fatalf("%s should be contract", addr)
	}
}
func TestFrom(t *testing.T) {
	addr := `0x8abfd268012d2113f31509bdd6e62d519c2ee621164d7687dc6ceb3eac8d55ec`
	conn, err := ethclient.Dial("https://api.myetherapi.com/eth")
	if err != nil {
		t.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	tx, _, err := conn.TransactionByHash(context.Background(), common.HexToHash(addr))
	t.Log(tx)
	txe := TransactionWithExtra{tx}
	if from := txe.From(); from.Hex() != common.HexToAddress("0xc5cf6410a3f2eda2b31ea73a4fd9b3e80d035fe1").Hex() {
		t.Fatalf("get from fail")
	}
	if contractAddr := txe.ContractAddress(); contractAddr.Hex() != common.HexToAddress("0x93e682107d1e9defb0b5ee701c71707a4b2e46bc").Hex() {
		t.Fatalf("get contract addr fail")
	}
}
