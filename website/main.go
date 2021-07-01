package main

import (
	"fmt"
	"html/template"
	"net/http"
)

const portNumber = ":7070"

// home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.html")
}

// about page handler
func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.html")
}

// this will help in rendering the templates
func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template", err)
		return
	}
}



func main() {
	
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)


	// to serve requests via browser
	_ = http.ListenAndServe(portNumber, nil)

	}

