package main

import (
	"log"
	"time"
)


type User struct {
	FirstName string
	LastName  string
	Age       int
	BirthDate time.Time
}

func main() {

	var myValue string
	myValue = "Green"

	log.Println("myValue is", myValue)
	changeUsingPointers(&myValue)
	log.Println("after using the function", myValue)

	user := User {
		FirstName: "Himanshu",
		LastName: "Yadav",
	}

	log.Println(user.FirstName)

}

func changeUsingPointers(s *string) {
	log.Println(s, *s)
	newValue := "Red"
	*s = newValue
}