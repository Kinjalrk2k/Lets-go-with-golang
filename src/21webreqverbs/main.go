package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to Web Verb Video")

	// PerformGetRequest()
	// PerformPostJsonRequest()
	PerformPostFormRequest()
}

func PerformGetRequest() {
	const myurl = "http://localhost:8000/get"

	response, err := http.Get((myurl))
	checkNillErr(err)
	defer response.Body.Close()

	fmt.Println("Status Code:", response.StatusCode)
	fmt.Println("Content Length:", response.ContentLength)

	content, _ := (io.ReadAll(response.Body))

	// fmt.Println(string(content))

	var responseString strings.Builder
	byteCount, _ := responseString.Write(content)
	fmt.Println("ByteCount is:", byteCount)
	fmt.Println(responseString.String())
}

func PerformPostJsonRequest() {
	const myurl = "http://localhost:8000/post"

	// fake json payload
	requestBody := strings.NewReader(`
		{
			"coursename":"Let's go with golang",
			"price": 0,
			"platform": "lco.in"
		}
	`)

	response, err := http.Post(myurl, "application/json", requestBody)
	checkNillErr(err)
	defer response.Body.Close()

	content, _ := (io.ReadAll(response.Body))

	fmt.Println(string(content))
}

func PerformPostFormRequest() {
	const myurl = "http://localhost:8000/postform"

	// form data
	data := url.Values{}
	data.Add("firstname", "Kinjal")
	data.Add("lastname", "Raykarmakar")
	data.Add("email", "kinjalrk2k@gmail.com")

	response, err := http.PostForm(myurl, data)
	checkNillErr(err)
	defer response.Body.Close()

	content, _ := (io.ReadAll(response.Body))

	fmt.Println(string(content))
}

func checkNillErr(err error) {
	if err != nil {
		panic(err)
	}
}
