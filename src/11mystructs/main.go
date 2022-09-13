package main

import "fmt"

func main() {
	fmt.Println("Structs in GoLang")

	kinjal := User{"Kinjal", "kinjalrk2k@gmail.com", true, 22}
	fmt.Println(kinjal)
	fmt.Printf("Details are: %+v\n", kinjal)

	kinjalAgain := User{Name: "Kinjal", Email: "kinjalrk2k@gmail.com", Status: true, Age: 22}
	fmt.Println(kinjalAgain)

	fmt.Printf("Name is %v and email is %v", kinjal.Name, kinjal.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
