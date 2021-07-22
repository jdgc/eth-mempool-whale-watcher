package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient/gethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

const (
	Wei   = 1
	GWei  = 1e9
	Ether = 1e18
)

func main() {
	nodeURL := os.Getenv("NODE_URL")

	if len(nodeURL) == 0 {
		log.Fatal("node url not set")
	}

	client, err := rpc.Dial(nodeURL)
	if err != nil {
		log.Fatal(err)
	}
	geth := gethclient.New(client)
	defer client.Close()

	logs := make(chan common.Hash)

	sub, err := geth.SubscribePendingTransactions(context.Background(), logs)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case tx := <-logs:
			go printTx(client, tx)
		}
	}
}

func printTx(client *rpc.Client, tx common.Hash) {
	var transaction map[string]interface{}
	client.Call(&transaction, "eth_getTransactionByHash", tx)

	// some transactions in pool return nothing when queried
	// we also want to ignore transactions with 0 eth value.
	if transaction == nil || transaction["value"] == "0x0" {
		return
	}

	if value, ok := transaction["value"].(string); ok {
		etherValue := valueInEth(value)

		switch comparison := etherValue.Cmp(big.NewFloat(5)); comparison {
		case 0:
			return
		case -1:
			return
		}

		fmt.Printf("*** NEW TX DETECTED ***\n")
		fmt.Printf("tx hash: %s\n", transaction["hash"])
		fmt.Printf("from: %s\n", transaction["from"])
		fmt.Printf("from: %s\n", transaction["to"])
		fmt.Printf("Ether value: %f\n", etherValue)
	} else {
		return
	}
}

func logBadValue(transaction map[string]interface{}) {
	fmt.Printf("bad value for tx hash: %s\n", transaction["hash"])
	fmt.Printf("value: %s\n", transaction["value"])
}

func valueInEth(hexValueInWei string) *big.Float {
	decodedValue, err := hexutil.DecodeBig(hexValueInWei)
	if err != nil {
		log.Fatalf("error decoding value: %s\n", hexValueInWei)
	}
	f := new(big.Float).SetInt(decodedValue)
	return new(big.Float).Quo(f, big.NewFloat(Ether))
}
