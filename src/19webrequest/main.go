package main

import (
	"fmt"
	"io"
	"net/http"
)

const url string = "https://lco.dev"

func main() {
	fmt.Println("LCO Web Request")

	response, err := http.Get(url)
	checkNillErr(err)
	defer response.Body.Close()

	fmt.Printf("Response is of type %T\n", response)

	dataBytes, err := io.ReadAll(response.Body)

	checkNillErr(err)
	content := string(dataBytes)
	fmt.Println("The content returned from the server is:\n", content)

}

func checkNillErr(err error) {
	if err != nil {
		panic(err)
	}
}
