package main

import "fmt"

type User struct {
	Name   string
	Email  string
	Age    int
	Status bool
}

func main() {
	fmt.Println("Welcome to struct session of GO...")

	pawan := User{"Pawan", "Pawan@go.dev", 30, true}
	fmt.Println("Details of User", pawan) //Details of User {Pawan Pawan@go.dev 30 true}

	//using +v give detailed info including teh variable name
	////User details {Name:Pawan Email:Pawan@go.dev Age:30 Status:true}
	fmt.Printf("User details %+v \n", pawan)

	fmt.Printf("User Name %v and Email %v", pawan.Name, pawan.Email) //User Name Pawan and Email Pawan@go.dev

}
