package utils

import (
	"fmt"
	"io/ioutil"
)

func ReadData() []byte {
	// Read plaintext
	filepath := "../Data/plaintext.txt"
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Printf("Failed to read the file: %v", filepath)
		panic(err)
	}

	return data
}
