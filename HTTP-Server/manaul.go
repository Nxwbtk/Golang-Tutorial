package main

import (
	"fmt"
	"log"
	"net/http"
)

func getHello(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Hello world")
}

func main () {
	http.HandleFunc("/hello", getHello)
	fmt.Println("Start at 8000")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}