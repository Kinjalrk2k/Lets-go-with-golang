package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=werjfjhsd8923u12ju23"

func main() {
	fmt.Println("Welcome to handling URLs in GoLang")

	fmt.Println(myurl)

	// parsing
	result, _ := url.Parse(myurl)

	fmt.Println("Scheme:", result.Scheme)
	fmt.Println("Host:", result.Host)
	fmt.Println("Path:", result.Path)
	fmt.Println("Port:", result.Port())
	fmt.Println("RawQuery:", result.RawQuery)

	qparams := result.Query()
	fmt.Printf("The type of query params is %T\n", qparams)
	fmt.Println("Course Name", qparams["coursename"])

	for _, val := range qparams {
		fmt.Println("Param is: ", val)
	}

	// need to pass reference
	partsOfUrl := &url.URL{
		Scheme:   "https",
		Host:     "lco.dev",
		Path:     "/tutcss",
		RawQuery: "user=kinjal",
	}

	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)

}
