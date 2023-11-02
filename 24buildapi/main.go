package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// model for Qlik Product - file
type QlikProduct struct {
	ProductId   string `json:"productid"`
	ProductName string `json:"productname"`
	Price       int    `json:"productprice"`
	Owner       *Owner `json:"owner"`
}

type Owner struct {
	Name    string `json:"ownername"`
	Website string `json:"website"`
}

// Fake database
var products []QlikProduct

// Check if QlikProduct is empty
// helper - file
func (p QlikProduct) isEmpty() bool {
	//return p.ProductId == "" && p.ProductName == ""
	return p.ProductName == ""
}

func main() {
	fmt.Println("API Learning - IRD")
	r := mux.NewRouter()

	//Seeding data
	products = append(products, QlikProduct{"1", "QlikView", 1000, &Owner{"QlikTech", "www.qlik.com"}})
	products = append(products, QlikProduct{"2", "QlikSense", 1500, &Owner{"QlikTech", "www.qlik.com"}})
	products = append(products, QlikProduct{"3", "DataIntegration", 1000, &Owner{"Talend", "www.talend.com"}})

	// Routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/products", getAllQlikProduct).Methods("GET")
	r.HandleFunc("/product/{id}", getQlikProduct).Methods("GET")
	r.HandleFunc("/product", createOneQlikProduct).Methods("POST")
	r.HandleFunc("/product/{id}", updateQlikProduct).Methods("PUT")
	r.HandleFunc("/product/{id}", deleteProductFromQlikProduct).Methods("DELETE")

	// Listening to the port
	log.Fatal(http.ListenAndServe(":3000", r))
}

//controller -file

// Serve home route
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API Learning by IRD</h1>"))
}

func getAllQlikProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all the list of Qlik Product")
	w.Header().Set("Conetent-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func getQlikProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get the Qlik Product")
	w.Header().Set("Conetent-Type", "application/json")

	//grab id from the request
	params := mux.Vars(r)

	//loop through the QlikProduct and find matching id and return thr response
	for _, product := range products {
		if product.ProductId == params["id"] {
			json.NewEncoder(w).Encode(product)
			return
		}
	}

	//Send a response when id is not found
	json.NewEncoder(w).Encode("No Product found with given ID")
}

func createOneQlikProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get the Qlik Product")
	w.Header().Set("Conetent-Type", "application/json")

	//what if body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	// what about -{}
	var product QlikProduct
	json.NewDecoder(r.Body).Decode(&product)
	if product.isEmpty() {
		json.NewEncoder(w).Encode("No product data inside JSON")
	}

	// Generate unique Product ID
	// append product into QlikProduct
	rand.Seed(time.Now().UnixNano())
	product.ProductId = strconv.Itoa(rand.Intn(100))

	products = append(products, product)
	json.NewEncoder(w).Encode(product)
	return
}

func updateQlikProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get the Qlik Product")
	w.Header().Set("Conetent-Type", "application/json")

	param := mux.Vars(r)

	for index, product := range products {
		if product.ProductId == param["id"] {

			//remove the Product entry form the QlikProduct slice
			products = append(products[:index], products[index+1:]...)

			//append the incoming product details to the QlikProduct slice
			var prod QlikProduct
			json.NewDecoder(r.Body).Decode(&prod)

			prod.ProductId = param["id"]
			products = append(products, prod)
			json.NewEncoder(w).Encode(prod)
			return
		}
	}

	//Send a response when id is not found
	json.NewEncoder(w).Encode("Product not found for the update")
}

func deleteProductFromQlikProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get the Qlik Product")
	w.Header().Set("Conetent-Type", "application/json")

	params := mux.Vars(r)

	for index, prod := range products {
		if prod.ProductId == params["id"] {
			products = append(products[:index], products[index+1:]...)
			json.NewEncoder(w).Encode(prod)
			return
		}
	}

	//Send a response when id is not found
	json.NewEncoder(w).Encode("Product not found for the update")
}
