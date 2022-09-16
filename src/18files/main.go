package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to files in golang")

	content := "This needs to go in a file - Kinjal Raykarmakar"

	file, err := os.Create("./mygofile.txt")
	checkNillErr(err)
	defer file.Close()

	length, err := io.WriteString(file, content)
	checkNillErr(err)

	fmt.Println("Length is:", length)

	readFile("./mygofile.txt")
}

func readFile(filename string) {
	dataByte, err := os.ReadFile(filename) // data is not a string, but in bytes
	checkNillErr(err)

	fmt.Println("Text data inside the file is\n", string(dataByte))
}

func checkNillErr(err error) {
	if err != nil {
		panic(err)
	}
}
