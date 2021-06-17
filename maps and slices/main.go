package main

import (
	"log"
	"sort"
)


type User struct {
	Firstname string
	Lastname string
}

func main() {
	myMap := make(map[string]string)
	myMap["dog"] = "jordan"
	log.Println(myMap["dog"])

	myUser := make(map[string]User)

	me := User {
		Firstname: "Himanshu",
		Lastname: "Yadav",
	}

	myUser["me"] = me
	log.Println(myUser["me"].Firstname)

	var mySlice []string
	mySlice = append(mySlice, "Himanshu")

	log.Println((mySlice))

	var mySlice2 []int

	mySlice2 = append(mySlice2, 2)
	mySlice2 = append(mySlice2, 1)
	mySlice2 = append(mySlice2, 3)

	log.Println(mySlice2)

	sort.Ints(mySlice2)

	log.Println(mySlice2)


}
