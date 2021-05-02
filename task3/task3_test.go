package task3

import (
	"bytes"
	"flag"
	"fmt"
	"math/big"
	"rsa_and_more/task1"
	"rsa_and_more/utils"
	"testing"
)

func TestPadding(t *testing.T) {
	// Read plaintext
	msg := utils.ReadData()

	fmt.Printf("\nOriginal msg: %s\n", msg)
	pad_msg := Pad512(msg)
	unpad_msg := Unpad512(pad_msg)
	fmt.Printf("\nUnpaded msg: %s\n", unpad_msg)

	if !bytes.Equal(msg, unpad_msg[:len(msg)]) {
		t.Error("The unpaded result is incorrect.")
	}
}

func TestOAEP(t *testing.T) {
	msg := utils.ReadData()
	size := flag.Int64("keysize", 1024, "specify key size")

	// Generate keys
	e, n, d := task1.GenerateKeys(*size)

	tryLimits := 10
	flag := false
	for i := 0; i < tryLimits; i++ {
		// Padding and Encrypt
		pad_msg := Pad512(msg)
		pad_msg_bi := new(big.Int)
		pad_msg_bi.SetBytes(pad_msg)
		ciphertext := task1.Encrypt(pad_msg_bi, e, n)

		// Decrypt and unpad
		decrypted_msg_bt := task1.Decrypt(ciphertext, d, n)
		decrypted_msg := decrypted_msg_bt.Bytes()
		unpad_msg := Unpad512(decrypted_msg)

		if bytes.Equal(msg, unpad_msg[:len(msg)]) {
			flag = true
		}
	}
	if !flag {
		t.Error("The unpaded result is incorrect.")
	}
}
