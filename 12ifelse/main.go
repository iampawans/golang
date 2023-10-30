package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Learn if else in GO...")

	// Code for user input
	fmt.Println("Please enter a number")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	//convert user input from string to int
	//inputNum, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	inputNum, err := strconv.Atoi(strings.TrimSpace(input))

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Input number is :", inputNum)
	}

	if inputNum%2 == 0 {
		fmt.Printf("Number %v is even\n", inputNum)
	} else {
		fmt.Printf("Number %v is odd\n", inputNum)
	}

	if inputNum > 0 && inputNum < 100 {
		fmt.Printf("Number %v is in range of 0-100\n", inputNum)
	}

}
