package main

func main() {
	// var card string = "Ace of Spades"
	// := <- Tells the compiler to figure out the type for variable assignment. (Type inferrment). Only used on instantiation of NEW variable.
	cards := newDeckFromFile("my_cards")
	cards.print()
	cards.shuffle()
	cards.print()
}
