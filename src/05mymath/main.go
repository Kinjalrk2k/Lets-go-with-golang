package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	fmt.Println("Math in Golang")

	// random numbers

	// math/crypto
	randomNumber, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(randomNumber)
}
