package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kinjalrk2k/mongoapi/router"
)

func main() {
	fmt.Println("MongoDB API")

	r := router.Router()

	fmt.Println("Server is getting started...")
	fmt.Println("Listening on PORT 4000")
	log.Fatal(http.ListenAndServe(":4000", r))
}
