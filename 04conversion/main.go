package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	welcome := "Welcome to Pizza World"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please neter your rating between 1 and 5")

	input, _ := reader.ReadString('\n')
	fmt.Println("Thank you for Rating : ", input)

	//strings is used for trim space in this case and string related operations
	num, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Converted int is : ", num+1)
	}

}
