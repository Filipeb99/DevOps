package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", HelloServer)
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		panic(err)
	}
}

func HelloServer(wrt http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(wrt, "Hello World!")
}
