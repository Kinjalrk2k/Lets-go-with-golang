package main

import "fmt"

func main() {
	fmt.Println("Methods in GoLang")

	kinjal := User{Name: "Kinjal", Email: "kinjalrk2k@gmail.com", Status: true, Age: 22}
	fmt.Println(kinjal)

	fmt.Printf("Name is %v and email is %v\n", kinjal.Name, kinjal.Email)
	kinjal.GetStatus()
	kinjal.NewMail()
	fmt.Println(kinjal.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active?", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@go.dev" // does not actually change the property
	fmt.Println("Email of this user is:", u.Email)
}
