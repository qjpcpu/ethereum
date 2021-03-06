package stats

import (
	"github.com/qjpcpu/log"
	"math/big"
	"testing"
)

func TestScan(t *testing.T) {
	log.SetLogLevel(log.DEBUG)
	scanner, err := GetScanner("/Users/jason/Library/Ethereum/geth.ipc", NewStatPrinter())
	//scanner, err := GetScanner("https://api.myetherapi.com/eth", NewStatPrinter())
	if err != nil {
		t.Fatal(err)
	}
	scanner.SubscribeAll()
	scanner.StartScan(big.NewInt(5270758), 1, 2)
}
