// main.go project main.go
package main

import (
	"fmt"
	"log"
	"net/http"
)

var page = "something in html"

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func pageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, page)
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/page", pageHandler)
	err := http.ListenAndServe("localhost:10000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
