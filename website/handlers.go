package main

import (
	"net/http"
)

// home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.html")
}

// about page handler
func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.html")
}
