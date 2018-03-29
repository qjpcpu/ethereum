package txview

import (
	"fmt"
	"github.com/qjpcpu/log"
	"math/big"
	"testing"
	"time"
)

func TestScan(t *testing.T) {
	log.SetLogLevel(log.DEBUG)
	rr := make(chan TxsInfo, 10)
	dd := make(chan TxResult, 10)
	scanner, err := GetScanner("http://localhost:18545", rr, dd)
	if err != nil {
		t.Fatal(err)
	}
	scanner.Subscribe("0x86fa049857e0209aa7d9e616f7eb3b3b78ecfdb0", "transfer(address,uint256)")
	go func() {
		scanner.StartScan(big.NewInt(5340682), 100, 2)
	}()
	go func() {
		time.Sleep(1 * time.Second)
		scanner.Close()
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
