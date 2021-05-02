package task1

import (
	"math/big"
	"rsa_and_more/utils"
)

func GenerateKeys(size int64) (pubKeyE *big.Int, pubKeyN *big.Int, priKeyD *big.Int) {
	e := big.NewInt(65537)
	n, phi := new(big.Int), new(big.Int)
	var p, q *big.Int
	for {
		p, q = utils.GenPrimePair(size)
		n.Mul(p, q)
		p.Sub(p, big.NewInt(1))
		q.Sub(q, big.NewInt(1))
		phi.Mul(p, q)
		temp := new(big.Int)
		temp.GCD(nil, nil, e, phi)
		if temp.Cmp(big.NewInt(1)) == 0 {
			break
		}
	}
	d := new(big.Int)
	d.ModInverse(e, phi)
	return e, n, d
}

func Encrypt(plainText *big.Int, pubKeyE *big.Int, pubKeyN *big.Int) *big.Int {
	cipherText := new(big.Int)
	cipherText.Exp(plainText, pubKeyE, pubKeyN)

	return cipherText
}

func Decrypt(cipherText *big.Int, priKeyD *big.Int, pubKeyN *big.Int) *big.Int {
	decryptResult := new(big.Int)
	decryptResult.Exp(cipherText, priKeyD, pubKeyN)

	return decryptResult
}
