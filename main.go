package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func handler(wtr http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	name := query.Get("name")
	key := query.Get("key")
	
	if name == "" {
		name = "Guest"
	}

	if key == "" {
		key = "0"
	}
	
	offset, err := strconv.Atoi(key)

	if err != nil {
		fmt.Fprintf(wtr, "Failed to convert key to int!\n%v", err)
		return
	}
	
	fmt.Fprintf(wtr, "Name : %s\n", name)
	fmt.Fprintf(wtr, "Key : %d\n\n", offset)

	offset %= 26

	runes := []rune(name)
	shift, maxSize := rune(offset), rune(26)

	for index, char := range runes {

		if char >= 'a'+shift && char <= 'z' ||
			char >= 'A'+shift && char <= 'Z' {

			char -= shift

		} else if char >= 'a' && char < 'a'+shift ||
			char >= 'A' && char < 'A'+shift {

			char += (maxSize - shift)

		}

		runes[index] = char
	}

	fmt.Fprintf(wtr, "Encrypted name : %s\n", string(runes))
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
