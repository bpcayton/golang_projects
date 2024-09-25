package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

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

	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

// Takes a deck of cards and a hand size. Returns one hand of specified size. Second return is remainder of deck.
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ", ")
}

func (d deck) saveToFile(fileName string) error {
	return os.WriteFile(fileName, []byte(d.toString()), 0666)
}

func newDeckFromFile(fileName string) deck {
	byteSlice, err := os.ReadFile(fileName)
	if err != nil {
		log.Println("Error: ", err)
		os.Exit(1)
	}
	// Convert the byteSlice into a string, split the string by the ',' delimiter, and convert to []string
	str := strings.Split(string(byteSlice), ", ")
	return deck(str)
}

// Loop through cards
// Generate a random number b/w 0 and length of slice
// Swap cards with the generated random number
func (d deck) shuffle() deck {

	// Seed the RNG
	source := rand.NewSource(time.Now().UnixNano())
	rando := rand.New(source)

	for i := range d {
		newPosition := rando.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
	return d
}
