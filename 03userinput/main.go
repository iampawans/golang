package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to Pizza rating"
	fmt.Println(welcome)

	// https://pkg.go.dev/bufio#Reader
	// Check for more information
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter the rating :")

	// comma OK / err
	// In case of expected err "input, err"
	// err will store the error
	// input is try and er is catch
	input, _ := reader.ReadString('\n') //\n indicates the end of input

	fmt.Println("Thank you for rating : ", input)
	fmt.Printf("Type of reader : %T \n", input)
}
