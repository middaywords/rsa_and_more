package task2

import (
	"math/big"
	"rsa_and_more/task1"
)

type AesServer struct {
	priKeyD *big.Int
	PubKeyE *big.Int
	PubKeyN *big.Int
}

func (s *AesServer) Setup(size int64) {
	s.PubKeyE, s.PubKeyN, s.priKeyD = task1.GenerateKeys(size)
}

// Receive messages in aes server from the client
func (s *AesServer) TestResponse(attemptKey_b *big.Int,
	wup []byte,
	C_i *big.Int,
) []byte {
	aesWrapper := AesCBCWrapper{
		Key: PaddingKey(attemptKey_b.Bytes()),
	}
	encryptedMsg := aesWrapper.CBC_encrypt(wup)

	response := s.Response(C_i, encryptedMsg)

	return response
}

func (s *AesServer) Response(C_i *big.Int,
	encryptedMsg []byte,
) []byte {
	// RSA decryption
	K_b := new(big.Int)
	temp := new(big.Int)
	temp.Exp(big.NewInt(2), big.NewInt(128), nil)
	// C_i^d % n % 2^128
	K_b.Exp(C_i, s.priKeyD, s.PubKeyN)
	K_b.Mod(K_b, temp)

	decAes := AesCBCWrapper{
		Key: PaddingKey(K_b.Bytes()),
	}
	decWup := decAes.CBC_decrypt(encryptedMsg)

	return decWup
}
