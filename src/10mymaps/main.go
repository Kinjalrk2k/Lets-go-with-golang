package main

import "fmt"

func main() {
	fmt.Println("Maps in GoLang")

	languages := make(map[string]string)
	languages["js"] = "JavaScript"
	languages["rb"] = "Ruby"
	languages["go"] = "Go"
	languages["py"] = "Python"

	fmt.Println("List of all languages:", languages)
	fmt.Println("JS is:", languages["js"])

	// deleting a value
	delete(languages, "rb")
	fmt.Println("List of all languages:", languages)

	// loop through maps
	for key, value := range languages { // can be for _, value := (comma-ok syntax)
		fmt.Printf("For key %v, value is %v\n", key, value)
	}
}
