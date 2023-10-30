package main

import "fmt"

func main() {
	fmt.Println("Welcome to loop learning in GO...")

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday"}

	//for loop 1
	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}

	//for loop 2
	// here day will behave as index but not value like c++
	for day := range days {
		fmt.Println(days[day])
	}

	//for loop 3
	//Here the first entry will be index and second entry will be the string
	//EX : The index is 0 and Day is Sunday
	for index, day := range days {
		fmt.Printf("The index is %v and Day is %v \n", index, day)
	}

	// for loop 4
	// Just like while loop
	index := 0
	for index < 10 {
		fmt.Println("index : ", index)
		index++ //++index in not valid in GO
	}

	// use of contiunue and break
	idx := 0
	for idx < 10 {

		if idx == 2 {
			idx++
			continue
		}

		if idx == 8 {
			break
		}

		fmt.Println("idx : ", idx)
		idx++
	}

	// use of GOTO
	ind := 0
	for ind < 10 {

		if ind == 5 {
			goto ird
		}

		fmt.Println("ind : ", ind)
		ind++ //++index in not valid in GO
	}

ird:
	fmt.Println("Pawan says HI")

}
