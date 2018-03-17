package stats

import (
	"github.com/qjpcpu/log"
	"math/big"
	"testing"
)

func TestScan(t *testing.T) {
	log.SetLogLevel(log.DEBUG)
	//scanner, err := GetScanner("/Users/jason/Library/Ethereum/geth.ipc", NewStatPrinter())
	scanner, err := GetScanner("https://api.myetherapi.com/eth", NewStatPrinter())
	if err != nil {
		t.Fatal(err)
	}
	//	scanner.SubscribeAll()
	scanner.Subscribe("0x0e0989b1f9b8a38983c2ba8053269ca62ec9b195")
	scanner.StartScan(big.NewInt(5268157), 1, 2)
}
