package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "http://lco.dev:3000/learn?coursename=reactJS&paymentid=irdpawan"

func main() {
	fmt.Println("handling URL in GO...")

	res, err := url.Parse(myUrl)

	if err != nil {
		panic(err)
	}

	fmt.Printf("res type %T\n", res)
	fmt.Println(res.Scheme)
	fmt.Println(res.Path)
	fmt.Println(res.Host)
	fmt.Println(res.Port())
	fmt.Println(res.RawQuery)

	pquery := res.Query()

	for key, value := range pquery {
		fmt.Printf("%v -> %v\n", key, value)
	}

	pQueryUrl := &url.URL{
		Scheme:  "http",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=pawan",
	}

	outQuery := pQueryUrl.String()

	fmt.Println(outQuery)
}
