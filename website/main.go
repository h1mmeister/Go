package main

import (
	"fmt"
	"net/http"
)


func main() {
	// this method takes an inline function with a response writer and a pointer to the request
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		n, err := fmt.Fprintf(w, "This is my first attempt to create a website in Go!")
		// check if there is any error
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(fmt.Sprintf("Number of bytes written: %d", n))
	})
	// to serve requests via browser
	_ = http.ListenAndServe(":7070", nil)

	}

