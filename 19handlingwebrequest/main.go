package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("Welcome to the learning of web request handling in GO...")

	url := "http://lco.dev"
	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Type of response %T \n", response)

	defer response.Body.Close()

	databyte, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	content := string(databyte)
	fmt.Println(content)

}
