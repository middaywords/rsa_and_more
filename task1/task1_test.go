package task1

import (
	"flag"
	"math/big"
	"rsa_and_more/utils"
	"testing"
)

func TestEncryptDecrypt(t *testing.T) {
	size := flag.Int64("keysize", 1024, "specify key size")

	// Generating keys
	e, n, d := GenerateKeys(*size)

	// Read plaintext
	dat := utils.ReadData()

	// Encryption
	plainText := new(big.Int)
	plainText.SetBytes(dat)
	cipherText := Encrypt(plainText, e, n)

	// Decryption
	decryptResult := Decrypt(cipherText, d, n)

	if decryptResult.Text(16) != plainText.Text(16) {
		t.Errorf(" The decryptiion result is incorrect. ")
	}
}
