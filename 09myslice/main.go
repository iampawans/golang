package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to Slice tutorial in GO..")

	//declartion od slice is similar to array
	var fruitSlice = []string{"Apple", "Orange", "banana"}
	fmt.Println("Fruit SLice after default", fruitSlice) //[Apple Orange banana]

	//Append is used to insert to the slice
	fruitSlice = append(fruitSlice, "mango", "Banana")
	fmt.Println("Fruit SLice after append", fruitSlice) //[Apple Orange banana mango Banana]

	// If you want ro print the range of the slice
	fruitSlice = append(fruitSlice[1:])
	fmt.Println("fruitSlice[1:]", fruitSlice) //[Orange banana mango Banana]

	fruitSlice = append(fruitSlice[1:3])
	fmt.Println("fruitSlice[1:3]", fruitSlice) //[banana mango]

	fruitSlice = append(fruitSlice[:3])
	fmt.Println("fruitSlice[:3]", fruitSlice) //[banana mango Banana]

	//Use of make is case of slice
	var numSlice = make([]int, 4)
	numSlice[0] = 123
	numSlice[1] = 234
	numSlice[2] = 345
	numSlice[3] = 456
	fmt.Println(numSlice)
	//check if slice is sorted
	fmt.Println(sort.IntsAreSorted(numSlice))

	//append creates new memory when we add new element to slice
	//that is the reason when 5th element is inserted to numSlice, it is success
	numSlice = append(numSlice, 89, 78, 67, 56)
	fmt.Println(numSlice)
	fmt.Println(sort.IntsAreSorted(numSlice))

	//sort the slices
	sort.Ints(numSlice)
	fmt.Println(numSlice)
	fmt.Println(sort.IntsAreSorted(numSlice))

	//how to delete any element at an index in slice
	var courses = []string{"c++", "java", "react", "javascript", "ruby", "go"}
	fmt.Println(courses) //[c++ java react javascript ruby go]
	index := 2

	// 0 to index and index + 1 to last
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses) //[c++ java javascript ruby go]

}
