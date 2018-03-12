package actdag

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
		ipcpath = "/root/.ethereum/geth.ipc"
	}
	if ethConn, err = ethclient.Dial(ipcpath); err != nil {
		log.Panicln(err)
	}
	return ethConn
}

func main() {
	client := getClient("")
	block, err := client.BlockByNumber(context.Background(), new(big.Int).SetInt64(100))
	log.Panicln(err, block)
}
