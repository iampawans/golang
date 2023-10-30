package main

import "fmt"

func main() {
	fmt.Println("Welcome to pointer learning in GO")

	var ptr *int
	fmt.Println("value of ptr: ", ptr)

	myNumber := 30
	ptr = &myNumber

	fmt.Println("new value of ptr: ", ptr)
	fmt.Println("new value of *ptr: ", *ptr)

	//perform some arithmatic operation on pointer
	*ptr = *ptr * 2
	fmt.Println("new value of ptr after operation: ", ptr)
	fmt.Println("new value of *ptr after operation: ", *ptr)
	fmt.Println("new value of myNumber after operation: ", myNumber)

}
