package txview

import (
	"fmt"
	"github.com/qjpcpu/log"
	"math/big"
	"testing"
)

func TestScan(t *testing.T) {
	log.SetLogLevel(log.DEBUG)
	rr := make(chan TxsInfo, 10)
	dd := make(chan TxResult, 10)
	scanner, err := GetScanner("https://api.myetherapi.com/eth", rr, dd)
	if err != nil {
		t.Fatal(err)
	}
	scanner.Subscribe("0x86fa049857e0209aa7d9e616f7eb3b3b78ecfdb0", "transfer(address,uint256)")
	go func() {
		scanner.StartScan(big.NewInt(5340682), 1, 2)
	}()
X:
	for {
		select {
		case msg := <-rr:
			fmt.Println(msg)
		case d := <-dd:
			fmt.Println("Exit:", d)
			break X
		}
	}
}
