package main

import "fmt"

func main() {
	fmt.Println("Welcome to array learning")

	// size must for defining array.
	// if no size given the error:= panic: runtime error: index out of range [0] with length 0
	var fruitArray [4]string

	fruitArray[0] = "Apple"
	fruitArray[1] = "Orange"
	fruitArray[3] = "Banana"

	//List of Fruits :  [Apple Orange  Banana]
	fmt.Println("List of Fruits : ", fruitArray)
	//List of Fruits :  4
	//len() will always return the actual aaray size
	fmt.Println("List of Fruits : ", len(fruitArray))
	fmt.Println("LFruits : ", fruitArray[1])

	// When array is declatred with predfined entries
	var vegArray = [3]string{"Potato", "Tomato", "Onion"}
	fmt.Println("List of Fruits : ", vegArray)

}
