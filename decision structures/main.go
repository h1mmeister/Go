package main

import "log"

func main() {

	myVar := "cat"

	switch myVar {
	case "cat":
		log.Println("That's a cat!")
	case "dog":
		log.Println("That's a dog!")
	default:
		log.Println("print nothing!")
	}
}