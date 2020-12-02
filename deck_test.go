package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	// code to make sure that a deck is created with a x number of cards
	d := newDeck()
	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	// code to make sure that the first card is an Ace of Spades
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be Ace of Spades but got %v", d[0])
	}

	// code to make sure that the last card is a Four of Clubs
	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card to be Four of Clubs but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckTestFromFile(t *testing.T) {
	os.Remove("_decktesting.txt")

	d := newDeck()

	d.saveToFile("_decktesting.txt")

	loadedDeck := newDeckFromFile("_decktesting.txt")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck, but got %v", len(loadedDeck))
	}

	os.Remove("_decktesting.txt")
}
