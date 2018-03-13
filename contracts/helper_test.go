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
	addr := `0x1ac95afe8008760b6c34221cffe0ea4b84e3a194712923c80c098d4c31ab0437`
	conn, err := ethclient.Dial("https://api.myetherapi.com/eth")
	if err != nil {
		t.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	tx, _, err := conn.TransactionByHash(context.Background(), common.HexToHash(addr))
	t.Log(tx)
	txe := TransactionWithExtra{tx}
	if from := txe.From(); from.Hex() != common.HexToAddress("0x2437b75ab0e43e2fe068c9833dc488723aa944ec").Hex() {
		t.Fatalf("get from fail")
	}
}
