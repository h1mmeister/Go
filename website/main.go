package main

import (
	"net/http"
)

const portNumber = ":7070"

func main() {
	
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	// to serve requests via browser
	_ = http.ListenAndServe(portNumber, nil)

	}

