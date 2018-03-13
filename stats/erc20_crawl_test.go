package stats

import (
	"github.com/qjpcpu/log"
	"math/big"
	"testing"
)

func TestScan(t *testing.T) {
	log.SetLogLevel(log.INFO)
	scanner, err := GetScanner("https://api.myetherapi.com/eth", `0x9a642d6b3368ddc662ca244badf32cda716005bc`, StatStoragePrinter{})
	if err != nil {
		t.Fatal(err)
	}
	scanner.StartScan(big.NewInt(5243713), 1)
}
