package main

import "fmt"

/**
* Create a new type of deck
* which is a clice of strings
 */
type deck []string

// Receiver (d deck) makes any variable of type: deck access to this method (print())
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

/**
* Create and return a list of cards (a new deck)
* Doesn't need a reciever since this is kind of acting like a c'tor
 */
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Hearts", "Aces", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}
