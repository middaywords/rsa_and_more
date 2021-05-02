package utils

import (
	"crypto/rand"
	"math/big"
)

func GenPrime(size int64) *big.Int {
	for {
		min := new(big.Int)
		min.Exp(big.NewInt(2), big.NewInt(size-1), nil)

		n, err := rand.Int(rand.Reader, min)
		if err != nil {
			panic("Error when generating a random number")
		}
		n.Add(n, min)
		// 20 times Miller Rabin test.
		if n.ProbablyPrime(20) {
			return n
		}
	}
}

func GenPrimePair(size int64) (*big.Int, *big.Int) {
	sizeP := size / 2
	sizeQ := size - sizeP
	pqOrFlag := false
	var p, q *big.Int
	for {
		p = GenPrime(sizeP)
		q = GenPrime(sizeQ)
		x := new(big.Int)
		x.Mul(p, q)
		k := int64(x.BitLen())
		if k == size {
			break
		} else if k > size && !pqOrFlag {
			sizeP -= 1
			pqOrFlag = true
		} else if k > size && pqOrFlag {
			sizeQ -= 1
			pqOrFlag = false
		} else if k < size && !pqOrFlag {
			sizeQ += 1
			pqOrFlag = true
		} else if k < size && pqOrFlag {
			sizeQ += 1
			pqOrFlag = false
		}
	}

	return p, q
}
