package main

import "fmt"

func main() {
	// var card string = "Ace of Spades"
	// := <- Tells the compiler to figure out the type for variable assignment. (Type inferrment). Only used on instantiation of NEW variable.
	cards := newDeck()
	for _, card := range cards {
		fmt.Println(card)
	}

	/**
	We don't necessarily need the index within the slice

	for index, card := range cards {
		fmt.Println(card)
	}
	*/
}

func newCard() string {
	return "5 of Diamonds"
}
