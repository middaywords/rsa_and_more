package utils

import (
	"crypto/rand"
	"errors"
	"math/big"
)

func GenRandNum(min, max *big.Int) (*big.Int, error) {
	// calculate the max we will be using
	bg := new(big.Int)
	if min.Cmp(max) >= 0 {
		return bg, errors.New("min >= max")
	}
	bg.Sub(max, min)
	n, err := rand.Int(rand.Reader, bg)
	if err != nil {
		panic(err)
	}
	n.Add(n, min)

	return n, nil
}
