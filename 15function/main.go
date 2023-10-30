package main

import "fmt"

// Main function is the entry point of the program
func main() {
	fmt.Println("Welcome to function session in GO...")
	greet()

	res := sum(1, 2)
	fmt.Println("Sum of 1 and 2 is : ", res)

	proRes := proSum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println("Sum of 1 to 10 is : ", proRes)

	res1, Sum1 := proTwoReturn("Hello", 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(res1)
	fmt.Println("Sum of 1 to 10 is : ", Sum1)

	a, b, c := proNReturn("Hello", 1, 2, 3, 4, 5)
	fmt.Println(a, b)
	fmt.Println("Sum of 1 to 5 is : ", c)
}

func greet() {
	fmt.Println("Hi from Go")
}

// when known parameter
func sum(varone int, varTwo int) int {
	return varone + varTwo
}

//Interesting : when n number of parameter
// can be valid for same type of parameter
func proSum(varOne ...int) int {
	total := 0

	for one := range varOne {
		total = total + one
	}
	return total
}

//Interesting : when n number of parameter
// when using ...int, this must be decclared in last
// there can be nay number of return type
func proTwoReturn(str string, varOne ...int) (string, int) {
	total := 0

	for one := range varOne {
		total = total + one
	}
	return str, total
}

func proNReturn(str string, varOne ...int) (string, string, int) {
	total := 0

	for one := range varOne {
		total = total + one
	}
	return str, " Buddy ", total
}
