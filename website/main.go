package main

import (
	"fmt"
	"net/http"
)

// home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	n, err := fmt.Fprintf(w, "This is my first attempt to create a website in Go! Here is the Home Page.")
	// check if there is any error
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(fmt.Sprintf("Number of bytes written: %d", n))
}

// about page handler
func About(w http.ResponseWriter, r *http.Request) {
	sum := aboutValues(1,1)
	_, _ = fmt.Fprintf(w, fmt.Sprintf("This is the %dnd page called About!", sum))

}

func aboutValues(x int, y int) int {
	return x+y
}


func main() {
	
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	// to serve requests via browser
	_ = http.ListenAndServe(":7070", nil)

	}

