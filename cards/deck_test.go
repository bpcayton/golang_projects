package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	const first_card string = "Ace of Spades"
	const last_card string = "King of Clubs"
	cards := newDeck()
	// Test for whole deck of cards
	if len(cards) != 52 {
		t.Errorf("Expected deck of 52 cards. Received deck of %d cards.", len(cards))
	}

	if cards[0] != first_card {
		t.Errorf("Expected first card to be '%s'. Received '%s'.", first_card, cards[0])
	}

	if cards[len(cards)-1] != last_card {
		t.Errorf("Expected last card to be '%s'. Received '%s'.", last_card, cards[len(cards)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	const FILENAME string = "_decktest"
	const LENGTH_OF_DECK = 52
	os.Remove(FILENAME)

	deck := newDeck()
	deck.saveToFile(FILENAME)

	loadedDeck := newDeckFromFile(FILENAME)
	if len(loadedDeck) != LENGTH_OF_DECK {
		t.Errorf("Expected %d cards in deck. Received %d", LENGTH_OF_DECK, len(loadedDeck))
	}
	os.Remove(FILENAME)

}
