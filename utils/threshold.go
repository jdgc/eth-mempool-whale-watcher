package utils

import (
	"log"
	"math/big"
	"os"
	"strconv"
)

const (
	DEFAULT  = 50
	BIT_SIZE = 64
)

var Threshold *big.Float

func LoadThreshold() {
	setting := os.Getenv("MONITOR_ETH_THRESHOLD")

	if setting == "" {
		log.Printf(RedString("No Threshold value supplied. Defaulting to %d ETH"), DEFAULT)
		Threshold = big.NewFloat(DEFAULT)
	} else {
		parsed, err := strconv.ParseFloat(setting, BIT_SIZE)
		if err != nil {
			log.Printf(RedString("Invalid Threshold value supplied. Defaulting to %d ETH"), DEFAULT)
			Threshold = big.NewFloat(DEFAULT)
			return
		}
		Threshold = big.NewFloat(parsed)
	}
}
