package main

import "log"

func main() {
	var word string
	word = saySomething("Himanshu")
	log.Println(word)
}

func saySomething(s string) string {
	return s
}