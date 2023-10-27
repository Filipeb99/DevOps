package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func formHandler(wtr http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/form" {
		http.Error(wtr, "404 not found", http.StatusNotFound)
		return
	}

	err := req.ParseForm()

	if err != nil {
		fmt.Fprintf(wtr, "ParseForm() error: %v", err)
		return
	}

	fmt.Fprintf(wtr, "POST request successful\n\n")

	name := req.FormValue("name")
	key, err := strconv.Atoi(req.FormValue("key"))

	if err != nil {
		fmt.Fprintf(wtr, "Failed to convert key to int!\n%v", err)
		return
	}

	fmt.Fprintf(wtr, "Name : %s\n", name)
	fmt.Fprintf(wtr, "Key : %d\n\n", key)

	key %= 26

	runes := []rune(name)
	offset, maxSize := rune(key), rune(26)

	for index, char := range runes {

		if char >= 'a'+offset && char <= 'z' ||
			char >= 'A'+offset && char <= 'Z' {

			char -= offset

		} else if char >= 'a' && char < 'a'+offset ||
			char >= 'A' && char < 'A'+offset {

			char += (maxSize - offset)

		}

		runes[index] = char
	}

	fmt.Fprintf(wtr, "Encrypted name : %s\n", string(runes))
}

func main() {
	fileServer := http.FileServer(http.Dir("."))

	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)

	fmt.Println("Starting server at port 8000")

	err := http.ListenAndServe(":8000", nil)

	if err != nil {
		log.Fatal(err)
	}
}
