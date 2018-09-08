package main

import (
	"math/big"
)

func weiToEth(wei *big.Int) float64 {
	x := new(big.Float).SetInt(wei)
	y := new(big.Float).SetInt(big.NewInt(10 ^ 18))
	res, _ := new(big.Float).Quo(x, y).Float64()
	return res
}
