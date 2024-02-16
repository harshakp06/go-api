package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Product struct {
	Id       string
	Name     string
	Quantity int
	Price    float64
}

var Products []Product

func main() {

	Products = []Product{
		{Id: "1", Name: "Chair", Quantity: 100, Price: 100.00},
		{Id: "2", Name: "Desk", Quantity: 200, Price: 200.00},
	}

	handleRequests()

}

func homepage(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Welcome to homepage")
	log.Println("Endpoint hit : hoempage")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/products", returnAllProducts)
	myRouter.HandleFunc("/", homepage)
	myRouter.HandleFunc("/product/{id}", getProduct)
	http.ListenAndServe("localhost:10000", myRouter)

	// http.HandleFunc("/products", returnAllProducts)
	// http.HandleFunc("/", homepage)
	// http.HandleFunc("/product/", getProduct)
	// http.ListenAndServe("localhost:10000", nil)

}

func returnAllProducts(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint hit: Homepage")
	json.NewEncoder(w).Encode(Products)

}

func getProduct(w http.ResponseWriter, r *http.Request) {
	// log.Println(r.URL.Path)
	vars := mux.Vars(r)
	key := vars["id"]
	// key := r.URL.Path[len("/product/"):]
	for _, product := range Products {
		if string(product.Id) == key {
			json.NewEncoder(w).Encode(product)
		}
	}

}
