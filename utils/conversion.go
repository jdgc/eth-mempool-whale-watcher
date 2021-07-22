package utils

import (
	"log"
	"math/big"

	"github.com/jdgc/eth-mempool-whale-watcher/constants"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

func ValueInEth(hexValueInWei string) *big.Float {
	decodedValue, err := hexutil.DecodeBig(hexValueInWei)
	if err != nil {
		log.Fatalf("error decoding value: %s\n", hexValueInWei)
	}
	f := new(big.Float).SetInt(decodedValue)
	return new(big.Float).Quo(f, big.NewFloat(constants.ETHER))
}

func ValueInGwei(hexValueInWei string) *big.Int {
	decodedValue, err := hexutil.DecodeBig(hexValueInWei)
	if err != nil {
		log.Fatalf("error decoding value: %s\n", hexValueInWei)
	}

	f := new(big.Float).SetInt(decodedValue)
	gweiFloat := new(big.Float).Quo(f, big.NewFloat(constants.GWEI))
	gweiInt, _ := gweiFloat.Int(nil)

	return gweiInt
}
