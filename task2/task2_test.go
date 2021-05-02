package task2

import (
	"bytes"
	"flag"
	"fmt"
	"rsa_and_more/utils"
	"testing"
)

func TestCCA2_attack(t *testing.T) {
	size := flag.Int64("keysize", 1024, "specify key size")

	// Initialize the server(generate aes keys)
	aesServer := AesServer{}
	aesServer.Setup(*size)

	// Initialize the attack(bind the server)
	attacker := AesAttacker{}
	attacker.SetServer(&aesServer)

	// Initialize the server(generate aes keys for a session)
	aesClient := AesClient{}
	aesClient.Setup(&aesServer)

	// Read message
	realMsg := utils.ReadData()

	// Client use the session key to encypt message
	cryptedMsg := aesClient.EncryptSendMsg(realMsg)

	// Attacker trying to guess the key
	guessKey := attacker.CCA2_attack(aesClient.EncryptedAesKey)
	fmt.Println("\nGuessed Key: ", guessKey.Text(16))

	// Any one can decrypt the message using the guessKey
	fmt.Println("\nTrying to decrypt the crypted message using the Guessed Key")
	aesWrapper := AesCBCWrapper{
		Key: guessKey.Bytes(),
	}
	decryptedMsg := aesWrapper.CBC_decrypt(cryptedMsg)

	fmt.Printf("\nReal msg: %s\n", realMsg)
	fmt.Printf("\nDecrypted msg: %s\n", decryptedMsg)
	if !bytes.Equal(PaddingBlocks(realMsg), decryptedMsg) {
		t.Error("The guessed key is incorrect.")
	}
}
