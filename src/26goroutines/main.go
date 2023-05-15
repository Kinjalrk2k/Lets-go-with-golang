package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var signals = []string{}

var wg sync.WaitGroup // a wait group; usually these are pointers!
var mut sync.Mutex    // usually a pointer

func main() {
	// go greeter("Hello") // fire up the routine, but we didn't wait
	// greeter("World")

	websites := []string{
		"http://lco.dev",
		"http://go.dev",
		"http://google.co.in",
		"http://fb.com",
		"http://github.com",
		"http://kinjal.dev",
	}

	for _, website := range websites {
		go getStatusCode(website) // starting go routines
		wg.Add(1)
	}

	wg.Wait() // waiting for the go routines to complete

	fmt.Println(signals)
}

// func greeter(s string) {
// 	for i := 0; i < 5; i++ {
// 		fmt.Println(s)
// 	}
// }

func getStatusCode(endpoint string) {
	defer wg.Done() // marking as done when done!

	response, err := http.Get(endpoint)
	if err != nil {
		log.Fatal(err)
	}

	// making this thread safe!
	mut.Lock()
	signals = append(signals, endpoint) // trying to access the same resource in multiple threads
	mut.Unlock()
	fmt.Printf("%s \t Status %d\n", endpoint, response.StatusCode)
}
