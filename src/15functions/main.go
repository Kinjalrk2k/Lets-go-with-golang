package main

import "fmt"

func main() {
	fmt.Println("Functions are Fun")

	greeter()

	result := adder(3, 5)
	fmt.Println("Result is", result)

	proRes, _ := proAdder(1, 2, -8, 5, 0, -2)
	fmt.Println("Pro Addiiton result is", proRes)
}

func greeter() {
	fmt.Println("Namastee from GoLang")
}

func adder(valOne int, valTwo int) int /* this is return type */ {
	return valOne + valTwo
}

func proAdder(values ...int) (int, string) {
	total := 0
	for _, value := range values {
		total += value
	}

	return total, "Hi from proAdder function"
}
