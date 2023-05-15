package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels")

	myChannel := make(chan int, 2) // buffered channel
	wg := &sync.WaitGroup{}

	wg.Add(2)

	// recieve only
	go func(ch <-chan int, wg *sync.WaitGroup) {
		val, isChannelOpen := <-myChannel
		fmt.Println(isChannelOpen)
		fmt.Println(val) // listens for values
		wg.Done()
	}(myChannel, wg)

	// send only
	go func(ch chan<- int, wg *sync.WaitGroup) {
		myChannel <- 5 // produces the value
		myChannel <- 6
		close(myChannel)
		wg.Done()
	}(myChannel, wg)

	wg.Wait()
}
