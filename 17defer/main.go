package main

import "fmt"

func main() {
	// defer input => Lets start, All, Hi ; Output=> Hi, ALl, Lets start
	//defer works on LIFO concept
	fmt.Println("Welcome to defer learning in GO...")
	// defer will execute in last
	defer fmt.Println("Lets start")
	defer fmt.Println("All")
	defer fmt.Println("Hi")

	fmt.Println("-----------------")

	// it will print 0-9 in reverse order because of defer
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

}
