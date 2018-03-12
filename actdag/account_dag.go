package main

import (
	"context"
	"log"
	"math/big"
	//"math/big"
	// "github.com/ethereum/go-ethereum/accounts/abi/bind"
	// "github.com/ethereum/go-ethereum/accounts/keystore"
	// "github.com/ethereum/go-ethereum/common"
	// "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func getClient(ipcpath string) *ethclient.Client {
	var ethConn *ethclient.Client
	var err error
	if ipcpath == "" {
		ipcpath = "/home/centos/.ethereum/testnet/geth.ipc"
	}
	if ethConn, err = ethclient.Dial(ipcpath); err != nil {
		log.Panicln(err)
	}
	return ethConn
}

func main() {
	client := getClient("")
	for bi := 0; bi < 10000; bi++ {
		block, err := client.BlockByNumber(context.Background(), new(big.Int).SetInt64(int64(bi)))
		if err != nil {
			log.Panicln(err)
		}
		txs := block.Transactions()
		log.Println("得到block:", block, "包含交易=", len(txs))
		for i, tx := range txs {
			log.Printf("%d. From=%v To=%v Val=%v [%v]\n", i, tx.To(), tx.To(), tx.Value(), tx.String())
		}
		if len(txs) > 0 {
			break
		}
	}
}
