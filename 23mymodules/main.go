package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Lets learn MOD in GO...")
	fmt.Println("Ref: Mod - https://go.dev/ref/mod")
	fmt.Println("Using Go Mod - https://go.dev/blog/using-go-modules")

	greeter()

	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")

	// running server on "http://localhost:4000/"
	// This is error prone func so use logs
	log.Fatal(http.ListenAndServe(":4000", r))
}

func greeter() {
	fmt.Println("Hey Mod Users")
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcone to golang series on MOD</h1>"))
}
