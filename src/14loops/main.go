package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops")

	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}

	// for loop
	for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
	}

	fmt.Println("----------------")

	for i := range days {
		fmt.Println(days[i])
	}

	fmt.Println("----------------")

	for index, day := range days {
		fmt.Printf("Index is %v and Day is %v\n", index, day)
	}

	fmt.Println("----------------")

	rougeValue := 1
	for rougeValue < 10 {

		if rougeValue == 5 {
			goto lco
		}

		fmt.Println("Value is:", rougeValue)
		rougeValue++
	}

lco:
	fmt.Println("Jumping...")
}
