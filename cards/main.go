package main

import "fmt"

func main() {
	// var card string = "Ace of Spades"
	// card := "Ace of Spades" // better way to assign variables
	// card := newCard()
	// card = "Five of Diamonds"

	cards := []string{"Ace of Spades", newCard()} // this is how we create a slice - need to mention the type as it will contain only one type
	cards = append(cards, "Six of Spades")        // this is how we append to the slice - it will create a new slice, will not modify it
	// fmt.Println(cards) // for printing

	for i, card := range cards { // this is how we iterate over a slice
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "Five of Diamonds"
}