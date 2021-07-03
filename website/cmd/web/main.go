package main

import (
	"github/h1mmeister/Go/pkg/handlers"
	"net/http"
)

const portNumber = ":7070"

func main() {
	
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	// to serve requests via browser
	_ = http.ListenAndServe(portNumber, nil)

	}

