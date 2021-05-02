package task2

import (
	"math/big"
	"rsa_and_more/utils"
)

type AesClient struct {
	EncryptedAesKey *big.Int
	aesKey          *big.Int
}

func (a *AesClient) Setup(server *AesServer) {
	min := new(big.Int)
	max := new(big.Int)
	min.Exp(big.NewInt(2), big.NewInt(127), nil)
	max.Exp(big.NewInt(2), big.NewInt(128), nil)
	aesKey, err := utils.GenRandNum(min, max)
	if err != nil {
		panic(err)
	}
	a.aesKey = aesKey

	a.EncryptedAesKey = new(big.Int)
	a.EncryptedAesKey.Exp(aesKey, server.PubKeyE, server.PubKeyN)
}

func (a *AesClient) EncryptSendMsg(msg []byte) []byte {
	aesWrapper := AesCBCWrapper{
		Key: a.aesKey.Bytes(),
	}
	cryptedMsg := aesWrapper.CBC_encrypt(msg)

	return cryptedMsg
}
