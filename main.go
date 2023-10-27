package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func handler(wtr http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(wtr, "Hello, %q", html.EscapeString(req.URL.Path))
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
