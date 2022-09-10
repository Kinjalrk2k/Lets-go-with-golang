package main

import "fmt"

func main() {
	fmt.Println("Welcome to a class on pointers")

	// var ptr *int // pointer that'll hold integer value
	// fmt.Println("Value of pointer is", ptr)

	myNumber := 23
	var ptr = &myNumber
	fmt.Println("Value of actual pointer is", ptr)  // 0x1400011c00
	fmt.Println("Value of actual pointer is", *ptr) // 23

	*ptr = *ptr + 2
	fmt.Println("New value is: ", myNumber) // 25
}
