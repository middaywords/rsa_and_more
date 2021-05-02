package task3

import (
	"crypto/sha512"
	"math/big"
	"rsa_and_more/utils"
)

const N = 1024
const K_0 = 512

func Pad512(msg []byte) []byte {
	r_hash := sha512.New()
	h_hash := sha512.New()

	// Generate a random number r
	max := new(big.Int)
	max.Exp(big.NewInt(2), big.NewInt(K_0), nil)
	r, err := utils.GenRandNum(big.NewInt(0), max)
	if err != nil {
		panic("Failed to generate random number.")
	}
	r_bytes := r.Bytes()

	// Padding zeros to r bytes
	r_bytes_K0 := make([]byte, K_0/8)
	for i := 0; i < len(r_bytes); i++ {
		r_bytes_K0[i] = r_bytes[i]
	}

	// Padding zeros to message
	msg_K0 := make([]byte, K_0/8)
	for i := 0; i < len(msg); i++ {
		msg_K0[i] = msg[i]
	}

	// msg XOR hash(random_number)
	r_hash.Write(r_bytes_K0)
	hash_res_r := r_hash.Sum(nil)
	result_x := make([]byte, K_0/8)
	for i := 0; i < K_0/8; i++ {
		result_x[i] = hash_res_r[i] ^ msg_K0[i]
	}

	// random_number XOR hash(X)
	h_hash.Write(result_x)
	hash_res_x := h_hash.Sum(nil)
	result_y := make([]byte, K_0/8)
	for i := 0; i < K_0/8; i++ {
		result_y[i] = hash_res_x[i] ^ r_bytes_K0[i]
	}

	result := append(result_x, result_y...)
	return result
}

func Unpad512(msg []byte) []byte {
	r_hash := sha512.New()
	h_hash := sha512.New()

	x := msg[:K_0/8]
	y := msg[K_0/8:]

	// y XOR hash(x)
	h_hash.Write(x)
	hash_res_x := h_hash.Sum(nil)
	result_r := make([]byte, K_0/8)
	for i := 0; i < K_0/8; i++ {
		result_r[i] = hash_res_x[i] ^ y[i]
	}

	// x XOR result_r
	r_hash.Write(result_r)
	hash_res_r := r_hash.Sum(nil)
	unpad_res := make([]byte, K_0/8)
	for i := 0; i < K_0/8; i++ {
		unpad_res[i] = hash_res_r[i] ^ x[i]
	}

	return unpad_res
}
