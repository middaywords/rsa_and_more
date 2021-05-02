package task2

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

type AesCBCWrapper struct {
	Key []byte
}

// Padding key to 32byte
func PaddingKey(key []byte) []byte {
	// Initially set to 0
	// Becuase it's hex byte array, so 128bit -> 32byte
	paddingBytes := make([]byte, 32-len(key))
	key = append(key, paddingBytes...)

	return key
}

// Padding plaintext to blocks
func PaddingBlocks(plaintext []byte) []byte {
	if len(plaintext) == 0 {
		plaintext = make([]byte, aes.BlockSize)
	}
	if len(plaintext)%aes.BlockSize != 0 {
		// initially 0
		paddingBytes := make([]byte, aes.BlockSize-len(plaintext)%aes.BlockSize)
		copy(paddingBytes, "")
		// Append slice.
		plaintext = append(plaintext, paddingBytes...)
	}

	return plaintext
}

func (a *AesCBCWrapper) CBC_encrypt(plaintext []byte) []byte {
	block, err := aes.NewCipher(a.Key)
	if err != nil {
		panic(err.Error())
	}
	plaintext = PaddingBlocks(plaintext)
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)

	return ciphertext
}

func (a *AesCBCWrapper) CBC_decrypt(ciphertext []byte) []byte {
	block, err := aes.NewCipher([]byte(a.Key))
	if err != nil {
		panic(err.Error())
	}
	if len(ciphertext) < aes.BlockSize {
		panic("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]
	if len(ciphertext)%aes.BlockSize != 0 {
		fmt.Println("cipher text len:", len(ciphertext))
		panic("ciphertext is not a multiple of the block size")
	}
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphertext, ciphertext)

	return ciphertext

}
