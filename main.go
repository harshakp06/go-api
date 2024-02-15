package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", homepage)
	http.ListenAndServe("localhost:10000", nil)

}

func homepage(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Welcome to homepage")
	fmt.Println("Endpoint hit : hoempage")
}
