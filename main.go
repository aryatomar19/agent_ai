package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "OK")
    })

    fmt.Println("hello world arya")
    http.ListenAndServe(":8080", nil)
}
