package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch case in GoLang")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Printf("Value of dice is: %v\n", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1 and you can open")
	case 2:
		fallthrough
	case 3:
		fallthrough
	case 4:
		fallthrough
	case 5:
		fmt.Printf("You can move %v spots\n", diceNumber)
	case 6:
		fmt.Println("You can move 6 spots and roll the dice again!")
	default:
		fmt.Println("What was that?!?")
	}
}
