package task2

import (
	"bytes"
	"fmt"
	"math/big"
)

type AesAttacker struct {
	server *AesServer
}

//
// Set the iter-th bit of guessKey to 1(val=true) or 0(val=false)
//
func (a *AesAttacker) getAttemptKey(guessKey *big.Int, iter int, val bool) *big.Int {
	temp := new(big.Int)
	attemptKey_b := new(big.Int)
	if val {
		// 1<<127
		temp.Exp(big.NewInt(2), big.NewInt(127), nil)
		// guessKey<<(127-iter)
		temp2 := new(big.Int)
		temp2.Exp(big.NewInt(2), big.NewInt(int64(127-iter)), nil).Mul(guessKey, temp2)
		// 1<<127 + guessKey<<(127-iter)
		attemptKey_b.Add(temp, temp2)
	} else {
		temp.Exp(big.NewInt(2), big.NewInt(int64(127-iter)), nil)
		attemptKey_b.Mul(temp, guessKey)
	}

	return attemptKey_b
}

func (a *AesAttacker) SetServer(pServer *AesServer) {
	a.server = pServer
}

func (a *AesAttacker) CCA2_attack(encryptedAesKey *big.Int) *big.Int {
	guessKey := big.NewInt(0)
	wup := []byte("test WUP request")
	for i := 0; i < 128; i++ {
		fmt.Printf("\n----------round %v----------\n", i)
		fmt.Println("WUP byte array:", wup)
		factor := new(big.Int)
		temp := new(big.Int)
		// temp = (127-i)*e
		temp.Mul(big.NewInt(int64(127-i)), a.server.PubKeyE)
		// factor = 2^((127-i)*e)
		factor.Exp(big.NewInt(2), temp, a.server.PubKeyN)
		C_i := new(big.Int)
		C_i.Mul(encryptedAesKey, factor).Mod(C_i, a.server.PubKeyN)
		fmt.Println("\nC_i: ", C_i.Text(16))

		// Set bit to 1 and attempt to decrypt
		attemptKey_b := a.getAttemptKey(guessKey, i, true)
		fmt.Println("Trying k", i, ": ", len(attemptKey_b.Text(16)), attemptKey_b.Text(16))

		response := a.server.TestResponse(attemptKey_b, wup, C_i)

		if bytes.Equal(response, wup) {
			// the current bit should be 1
			temp2 := new(big.Int)
			temp2.Exp(big.NewInt(2), big.NewInt(int64(i)), nil)
			guessKey.Add(temp2, guessKey)
			temp3 := new(big.Int)
			temp3.Exp(big.NewInt(2), big.NewInt(int64(127-i)), nil)
			temp3.Mul(guessKey, temp3)
			fmt.Println("k", 127-i, ":", temp3.Text(16))
		} else {
			// the current bit should be 0, resend the forged message
			attemptKey_b = a.getAttemptKey(guessKey, i, false)
			fmt.Println("Trying k", i, ": ", len(attemptKey_b.Text(16)), attemptKey_b.Text(16))
			response = a.server.TestResponse(attemptKey_b, wup, C_i)
			fmt.Println("\nResponse: ", response)

			if bytes.Equal(response, wup) {
				fmt.Println("k", 127-i, ":", guessKey.Text(16))
			} else {
				panic("No matching bit.")
			}
		}

		fmt.Println("\n--------------------------")
	}

	return guessKey
}
