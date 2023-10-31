package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func main() {
	fmt.Println("### Welcome to web requests GET/POST/POSTFORM in GO ###")

	fmt.Println("Hey, Please enter the Web request Type(GET/POST/POSTFORM)")
	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')
	CheckErrNil(err)

	inputString := strings.TrimSpace(strings.ToUpper(input))
	fmt.Println("You have entered:", inputString)

	switch inputString {
	case "GET":
		GetRequest()
	case "POST":
		PostRequest()
	case "POSTFORM":
		PostFormRequest()
	default:
		fmt.Println("We are supporting GET/POST/POSTFORM")
	}

}

func GetRequest() {
	const myUrl = "http://localhost:8000/get"

	response, err := http.Get(myUrl)
	CheckErrNil(err)

	//This is MUST for every http call
	defer response.Body.Close()

	byteData, err1 := io.ReadAll(response.Body)
	CheckErrNil(err1)

	//Type 1 :Read Byte to string
	//fmt.Println(string(byteData))

	//Type 2 : Read using strings builder
	var responseString strings.Builder
	contentLen, err2 := responseString.Write(byteData)
	CheckErrNil(err2)
	fmt.Println("Conetent length is:", contentLen)
	fmt.Println(responseString.String())

}

func PostRequest() {
	const myUrl = "http://localhost:8000/post"

	//fake json payload
	body := strings.NewReader(`
		{
			"name":"Pawan Kumar",
			"trigram":"ird"
		}
	`)

	response, err := http.Post(myUrl, "application/json", body)
	CheckErrNil(err)

	defer response.Body.Close()

	dataByte, err := io.ReadAll(response.Body)
	CheckErrNil(err)

	var responseString strings.Builder
	contentLen, err := responseString.Write(dataByte)
	CheckErrNil(err)
	fmt.Println("Content Length :", contentLen)
	fmt.Println(responseString.String())
}

func PostFormRequest() {
	const myUrl = "http://localhost:8000/postform"

	//form data
	data := url.Values{}
	data.Add("name", "Pawan Kumar")
	data.Add("Age", "30")
	data.Add("country", "India")

	response, err := http.PostForm(myUrl, data)
	CheckErrNil(err)

	defer response.Body.Close()

	dataByte, err := io.ReadAll(response.Body)
	CheckErrNil(err)

	fmt.Println(string(dataByte))
}

func CheckErrNil(err error) {
	if err != nil {
		panic(err)
	}
}
