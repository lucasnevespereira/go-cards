package main

import (
	"fmt"
	"strings"
)

// Create a new type of deck
// which is a slice of strings
type deck []string

// returning list of cards(deck)
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, suit+" of "+value)
		}
	}

	return cards
}

// Receiver function
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// deal func with multiple returning values
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	// deck is a slice of strings []string
	// let's join a slice of strings together
	return strings.Join([]string(d), ",")
}
