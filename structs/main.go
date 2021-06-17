package main

import "log"

type myStruct struct {
	Firstname string
}
// Recievers
func (m *myStruct) printFirstName() string {
	return m.Firstname
}


func main() {

	var myVar myStruct
	myVar.Firstname = "Himanshu"

	myVar2 := myStruct{
		Firstname: "Simran",
	}

	log.Println("myVar is set to", myVar.printFirstName())
	log.Println("myVar2 is set to", myVar2.printFirstName())

}