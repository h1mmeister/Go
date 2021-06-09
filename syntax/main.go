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

	user := User {
		FirstName: "Himanshu",
		LastName: "Yadav",
	}

	log.Println(user.FirstName)

}
