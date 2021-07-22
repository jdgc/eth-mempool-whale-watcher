package utils

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

func TestValueInEth(t *testing.T) {
	wei := hexutil.EncodeBig(big.NewInt(1000000000000000000))
	result := ValueInEth(wei)
	expected := big.NewFloat(1)

	if result.Cmp(expected) != 0 {
		t.Errorf("Expected %f but got %f", expected, result)
	}

	wei = hexutil.EncodeBig(big.NewInt(0))
	result = ValueInEth(wei)
	expected = big.NewFloat(0)

	if result.Cmp(expected) != 0 {
		t.Errorf("Expected %f but got %f", expected, result)
	}
}

func TestValueInGwei(t *testing.T) {
	wei := hexutil.EncodeBig(big.NewInt(100000000000000))
	result := ValueInGwei(wei)
	expected := big.NewInt(100000)

	if result.Cmp(expected) != 0 {
		t.Errorf("Expected %d but got %d", expected, result)
	}

	wei = hexutil.EncodeBig(big.NewInt(0))
	result = ValueInGwei(wei)
	expected = big.NewInt(0)

	if result.Cmp(expected) != 0 {
		t.Errorf("Expected %d but got %d", expected, result)
	}
}
