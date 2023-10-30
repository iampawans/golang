package main

import (
	"fmt"
)

func main() {

	fmt.Println("Welcome to map in GO...")

	var language = make(map[string]string)

	language["GO"] = "Golang"
	language["JS"] = "JavaScript"
	language["RB"] = "Ruby"
	language["PY"] = "Python"

	// the output is laways sorted
	fmt.Println("Laguages are : ", language)                //Laguages are :  map[GO:Golang JS:JavaScript PY:Python RB:Ruby]
	fmt.Println("Laguages GO represent : ", language["GO"]) //Laguages GO represent :  Golang

	fmt.Println("Size of language : ", len(language)) //Size of language :  4

	//loop in map
	for key, value := range language {
		fmt.Printf("Key is %v and Value is %v \n", key, value)
	}

	for _, value := range language {
		fmt.Printf("language is %v \n", value)
	}

	//remove elment in map
	delete(language, "PY")
	fmt.Println("Laguages are : ", language) //Laguages are :  map[GO:Golang JS:JavaScript RB:Ruby]

}
