package main

import "fmt"

func main() {
	fmt.Println("Welcome to methods in GO...")

	pawan := User{"Pawan", "ird@go.dev", true, 30, 9945679209}
	fmt.Println("User details :", pawan)
	fmt.Println("Email Id using method:", pawan.GetEmail())
	pawan.SetEmail("test@go.dev")
	fmt.Println("Check Email Id using method:", pawan.GetEmail())
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
	mobile int
}

//method to get email id
func (u User) GetEmail() string {
	return u.Email
}

//method to set email id
func (u User) SetEmail(email string) {
	u.Email = email
	fmt.Println(u.Email)
}
