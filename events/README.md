# example

```
package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/qjpcpu/ethereum/events"
	"time"
)

func main() {
	var block_number uint64 = 3127385
	abi_str := `[{"constant":true,"inputs":[],"name":"totalSupply","outputs":[{"name":"supply","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"anonymous":false,"inputs":[{"indexed":true,"name":"from","type":"address"},{"indexed":true,"name":"to","type":"address"},{"indexed":false,"name":"value","type":"uint256"}],"name":"Transfer","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"name":"owner","type":"address"},{"indexed":false,"name":"spender","type":"address"},{"indexed":false,"name":"value","type":"uint256"}],"name":"Approval","type":"event"}]`
	contractAddr := `0x08323a38ef99102a1d02c182e7c86c57994ae0d8`
	conn, _ := ethclient.Dial(`http://localhost:18545`)
	dataCh, errCh := make(chan events.Event, 1), make(chan error, 1)
	b := events.NewScanBuilder()
	rep, err := b.SetClient(conn).
		SetEvents("Transfer").
		SetContract(common.HexToAddress(contractAddr)).
		SetABI(abi_str).
		SetFrom(block_number).
		SetGracefullExit(true).
		BuildAndRun(dataCh, errCh)
	if err != nil {
		fmt.Println(err)
		return
	}
	done := rep.WaitChan()
	go func() {
		time.Sleep(10 * time.Second)
		rep.Stop()
	}()
	for {
		select {
		case d := <-dataCh:
			fmt.Printf("bn:%v tx:%s contract:%s %s:%s\n", d.BlockNumber, d.TxHash.Hex(), d.Address.Hex(), d.Name, d.Data.String())
		case err = <-errCh:
			fmt.Println("err:", err)
		case <-done:
			fmt.Println("exit")
			return
		}
	}
}
```
