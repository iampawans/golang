package main

import "fmt"

func main() {
	var name string = "Pawan"
	fmt.Println(name)
	fmt.Printf("Variable type is %T \n", name)

	var number uint8 = 255
	fmt.Println(number)
	fmt.Printf("Variable type is %T \n", number)

	var number1 int = 256
	fmt.Println(number1)
	fmt.Printf("Variable type is %T \n", number1)

	var floatNumber float32 = 255.4576589387637739
	fmt.Println(floatNumber)
	fmt.Printf("Variable type is %T \n", floatNumber)

	var floatNumber1 float64 = 255.4576589387637739
	fmt.Println(floatNumber1)
	fmt.Printf("Variable type is %T \n", floatNumber1)

	// data type is not required, Var auto detects
	var number2 = 255
	fmt.Println(number2)
	fmt.Printf("Variable type is %T \n", number2)

	// only name can be used for data type using :=
	// This can be only used inside function
	name1 := "Evika"
	fmt.Println(name1)
	fmt.Printf("Variable type is %T \n", name1)
}
