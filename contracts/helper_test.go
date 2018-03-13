package contracts

import (
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
