package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("### Welcome to Encode and Decode JSOn in GO ###")
	EncodeJson()
	DecodeJson()
}

type car struct {
	Name       string   `json:"CarName"`
	Type       []string `json:"CarType,omitempty"` //omitempty will ognore the entries with value as nil
	isElectric bool     `json:"IsElectric"`        //bool is not exported
	Price      int      `json:"CarPrice"`
	ChesisNum  string   `json:"-"` // This can be used when you want to hide the entry. ex: Password
}

// Encode Json
func EncodeJson() {
	carDetails := []car{
		{"Hyundai Creta", []string{"SUV", "MUV"}, true, 2000000, "AWHG765GFR96Q1"},
		{"TATA Nexon", []string{"SUV", "MUV"}, true, 1500000, "FR96Q1AWHG765G"},
		{"Honda Civic", []string{"Sedan"}, false, 1800000, "765GFAWHGR96Q1"},
		{"Jeep Compass", nil, false, 2500000, "Q176AW5GFR96HG"},
	}

	//Package the data as Json data

	/*
		//OUTPUT 1:
		//[{"Name":"Hyundai Creta","Type":["SUV","MUV"],"Price":2000000,"ChesisNum":"AWHG765GFR96Q1"},
		//{"Name":"TATA Nexon","Type":["SUV","MUV"],"Price":1500000,"ChesisNum":"FR96Q1AWHG765G"},
		//{"Name":"Honda Civic","Type":["Sedan"],"Price":1800000,"ChesisNum":"765GFAWHGR96Q1"},
		//{"Name":"Jeep Compass","Type":["SUV"],"Price":2500000,"ChesisNum":"Q176AW5GFR96HG"}]
		data, err := json.Marshal(carDetails)
		if err != nil {
			panic(err)
		}

		fmt.Println(string(data))
	*/

	//OUTPUT 2: func json.MarshalIndent(v any, prefix string, indent string) ([]byte, error)
	/*
		type car struct {
			Name       string
			Type       []string
			isElectric bool
			Price      int
			ChesisNum  string
		}
		[
				{
						"Name": "Hyundai Creta",
						"Type": [
								"SUV",
								"MUV"
						],
						"Price": 2000000,
						"ChesisNum": "AWHG765GFR96Q1"
				},
		]
	*/

	//OUTPUT 3 : When Struct has Json keyword appended
	/*
		type car struct {
			Name       string   `json:"CarName"`
			Type       []string `json:"CarType,omitempty"`
			isElectric bool
			Price      int    `json:"CarPrice"`
			ChesisNum  string `json:"-"`
		}

		[
				{
						"CarName": "Hyundai Creta",
						"CarType": [
								"SUV",
								"MUV"
						],
						"CarPrice": 2000000
				},
				...
				{
						"CarName": "Jeep Compass",
						"CarPrice": 2500000
				}
		]
	*/
	data, err := json.MarshalIndent(carDetails, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))

}

func DecodeJson() {

	var carDetails car
	//Web tresponse will always be in bytes
	jsonDatafromWeb := []byte(`
		{
			"CarName": "Hyundai Creta",
			"CarType": ["SUV","MUV"],
			"CarPrice": 20000
		}
	`)

	valiadteJson := json.Valid(jsonDatafromWeb)

	if valiadteJson {
		fmt.Println("Valid JSON")

		//Convert Byte data to JSON data
		json.Unmarshal(jsonDatafromWeb, &carDetails)
		fmt.Println(carDetails)         //{Hyundai Creta [SUV MUV] false 2000000 }
		fmt.Printf("%#v\n", carDetails) // main.car{Name:"Hyundai Creta", Type:[]string{"SUV", "MUV"}, isElectric:false, Price:2000000, ChesisNum:""}
	} else {
		fmt.Println("Invalid JSON")
	}

	// In SOme cases, you might want to store data in key. value pair
	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDatafromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData) // map[string]interface {}{"CarName":"Hyundai Creta", "CarPrice":2e+06, "CarType":[]interface {}{"SUV", "MUV"}}

	//If you want to iterate key value
	for key, value := range myOnlineData {
		fmt.Printf("Key: %v and Value: %v and Type: %T\n", key, value, value)
	}

}
