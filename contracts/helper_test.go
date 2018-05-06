package contracts

import (
	"context"
	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
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
	addr := `0x8abfd268012d2113f31509bdd6e62d519c2ee621164d7687dc6ceb3eac8d55ec`
	conn, err := ethclient.Dial("https://api.myetherapi.com/eth")
	if err != nil {
		t.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	tx, _, err := conn.TransactionByHash(context.Background(), common.HexToHash(addr))
	t.Log(tx)
	txe := TransactionWithExtra{tx}
	if from := txe.From(); from.Hex() != common.HexToAddress("0xc5cf6410a3f2eda2b31ea73a4fd9b3e80d035fe1").Hex() {
		t.Fatalf("get from fail")
	}
	if contractAddr := txe.ContractAddress(); contractAddr.Hex() != common.HexToAddress("0x93e682107d1e9defb0b5ee701c71707a4b2e46bc").Hex() {
		t.Fatalf("get contract addr fail")
	}
}

func TestStatus(t *testing.T) {
	addr := `0xbea29cce7780090fe6e8fa4ce16acd9a684d6a8b931a422dfa14cd66370836ec`
	conn, err := ethclient.Dial("https://api.myetherapi.com/eth")
	if err != nil {
		t.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	tx, _, err := conn.TransactionByHash(context.Background(), common.HexToHash(addr))
	txe := TransactionWithExtra{tx}
	valid, err := txe.IsSuccess(conn)
	if err != nil {
		t.Fatal(err)
	}
	if valid {
		t.Fatal("shold invalid")
	}
	addr = `0x78ef04aede619ed9395bb2b2bde12d6a2320d2d54d8db4522a7a65f400f8d427`
	tx, _, _ = conn.TransactionByHash(context.Background(), common.HexToHash(addr))
	txe = TransactionWithExtra{tx}
	valid, err = txe.IsSuccess(conn)
	if err != nil {
		t.Fatal(err)
	}
	if !valid {
		t.Fatal("shold valid")
	}
}

func TestDeployContract(t *testing.T) {
	conn, err := ethclient.Dial("/home/ubuntu/repository/eth_home/data0/geth.ipc")
	if err != nil {
		t.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	keyJson := `{"address":"892d4394ff96164eab6121fa54657190dd37987c","crypto":{"cipher":"aes-128-ctr","ciphertext":"876b46fbe758221bbdae2e8d47eca32fbfabcfd6d5de2f2eb7e544363dcb50bb","cipherparams":{"iv":"e5c093fdad4d5d24405f363b6e93c7f8"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"ca1b094454bb6b4972f03b3011f0e25198ef1bed7dacd2168757c5940bfcfcc4"},"mac":"35c3c3fd60fdfba39391129c8486f5290709234a3c8295b6a164285465972c9b"},"id":"c90b1923-930f-4160-98b1-3febbbe03ac2","version":3}`
	keyPasswd := "123"
	tABI := `[{"constant":true,"inputs":[],"name":"name","outputs":[{"name":"","type":"string"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"_spender","type":"address"},{"name":"_value","type":"uint256"}],"name":"approve","outputs":[{"name":"success","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"totalSupply","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"_from","type":"address"},{"name":"_to","type":"address"},{"name":"_value","type":"uint256"}],"name":"transferFrom","outputs":[{"name":"success","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[],"name":"decimals","outputs":[{"name":"","type":"uint8"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"version","outputs":[{"name":"","type":"string"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[{"name":"_owner","type":"address"}],"name":"balanceOf","outputs":[{"name":"balance","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":true,"inputs":[],"name":"symbol","outputs":[{"name":"","type":"string"}],"payable":false,"stateMutability":"view","type":"function"},{"constant":false,"inputs":[{"name":"_to","type":"address"},{"name":"_value","type":"uint256"}],"name":"transfer","outputs":[{"name":"success","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":false,"inputs":[{"name":"_spender","type":"address"},{"name":"_value","type":"uint256"},{"name":"_extraData","type":"bytes"}],"name":"approveAndCall","outputs":[{"name":"success","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[{"name":"_owner","type":"address"},{"name":"_spender","type":"address"}],"name":"allowance","outputs":[{"name":"remaining","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"},{"inputs":[{"name":"_initialAmount","type":"uint256"},{"name":"_tokenName","type":"string"},{"name":"_decimalUnits","type":"uint8"},{"name":"_tokenSymbol","type":"string"}],"payable":false,"stateMutability":"nonpayable","type":"constructor"},{"payable":false,"stateMutability":"nonpayable","type":"fallback"},{"anonymous":false,"inputs":[{"indexed":true,"name":"_from","type":"address"},{"indexed":true,"name":"_to","type":"address"},{"indexed":false,"name":"_value","type":"uint256"}],"name":"Transfer","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"name":"_owner","type":"address"},{"indexed":true,"name":"_spender","type":"address"},{"indexed":false,"name":"_value","type":"uint256"}],"name":"Approval","type":"event"}]`
	tBIN := "0x"

	addr, tx, err := DeployContract(conn, keyJson, keyPasswd, tABI, tBIN, big.NewInt(100), "i-o-p", uint8(18), "iOP")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(addr.String(), tx)
}

func TestX(t *testing.T) {
	t.Log(Keccak256Hash("Birth(address,uint256,uint256,uint256,uint256)").Hex())
	topic := Keccak256Hash("Birth(address,uint256,uint256,uint256,uint256)")
	conn, err := ethclient.Dial("/Users/jason/repository/private-chain/data0/geth.ipc")
	if err != nil {
		t.Fatal(err)
	}
	ch := make(chan types.Log)
	fq := ethereum.FilterQuery{
		FromBlock: big.NewInt(5342459),
		//		ToBlock:   big.NewInt(5342460),
		Addresses: []common.Address{common.HexToAddress("0x06012c8cf97bead5deae237070f9587f8e7a266d")},
		Topics:    [][]common.Hash{[]common.Hash{topic}},
	}
	if _, err = conn.SubscribeFilterLogs(context.Background(), fq, ch); err != nil {
		t.Fatal(err)
	}
	logs, err := conn.FilterLogs(context.TODO(), ethereum.FilterQuery{
		FromBlock: big.NewInt(6342459),
		//		ToBlock:   big.NewInt(5342460),
		Addresses: []common.Address{common.HexToAddress("0x06012c8cf97bead5deae237070f9587f8e7a266d")},
		Topics:    [][]common.Hash{[]common.Hash{topic}},
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(len(logs))
	// l1 := logs[0]
	// t.Log(l1.String())
	// t.Log(common.BytesToAddress(l1.Data[:32]).Hex())
	// t.Log(new(big.Int).SetBytes(l1.Data[32:64]).String())
	// t.Log(new(big.Int).SetBytes(l1.Data[64:96]).String())
	// t.Log(new(big.Int).SetBytes(l1.Data[96:128]).String())
}

func TestWaitTx(t *testing.T) {
	addr := `0x33c5a13303945926cff78678803c499fa2b13410ec150b5feccc423987be17b4` // EOS contract
	conn, err := ethclient.Dial("http://10.140.0.4:8545")
	if err != nil {
		t.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	rep, err := conn.TransactionReceipt(context.Background(), common.HexToHash(addr))
	t.Log(rep, err)
}
