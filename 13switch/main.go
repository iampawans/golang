package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to the switch learning...")

	// Seed the random number generator with the current time
	rand.Seed(time.Now().UnixNano())

	// Generate a random number between 1 and 6
	randomNum := rand.Intn(6) + 1

	fmt.Printf("Number generate is : %v\n", randomNum)

	switch randomNum {
	case 1:
		fmt.Println("Open the game with Number 1")
	case 2:
		fmt.Println("Move to space 2")
	case 3:
		fmt.Println("Move to space 3")
	case 4:
		fmt.Println("Move to space 4")
		fallthrough
	case 5:
		fmt.Println("Move to space 5")
	case 6:
		fmt.Println("Move to space 6 and generate num again")
	default:
		fmt.Println("Invalid numver generated")
	}
}
