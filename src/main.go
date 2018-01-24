// main.go project main.go
package main

import (
    "fmt"
    "log"
    "net/http"
)

func h(w http.ResponseWriter, r *http.Request) {
    http.Header.Add(w.Header(), "Content-Type", "text/json; charset=utf-8")
    fmt.Fprintf(w, `{"message":"Hello World!"}`)
}

func main() {
    http.HandleFunc("/helloworld", h)
    e := http.ListenAndServe("0.0.0.0:5000", nil)
    if e != nil {
        log.Fatal("ListenAndServe: ", e)
    }
}
