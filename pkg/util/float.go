package util

import "math/big"

func NewFloat() *big.Float {
	result := big.NewFloat(0)
	result.SetPrec(256)
	return result
}
