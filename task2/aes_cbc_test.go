package task2

import (
	"bytes"
	"fmt"
	"testing"
)

func TestCBC(t *testing.T) {
	key := []byte("6368616e676520746869732070617373")
	plaintext := []byte("exampleplaintext2")

	aesWrap := AesCBCWrapper{
		Key: key,
	}
	ciphertext := aesWrap.CBC_encrypt(plaintext)
	fmt.Printf("%x\n", ciphertext)
	decRes := aesWrap.CBC_decrypt(ciphertext)
	fmt.Printf("%s\n", decRes)

	if !bytes.Equal(decRes, PaddingBlocks(plaintext)) {
		t.Error("The decryption result is not correct")
	}
}
