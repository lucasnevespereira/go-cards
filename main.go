// Reminder: Package main defines an executable output
package main

import "fmt"

func main() {

	// var card string = "Ace of Spades"
	// or
	// card := "Ace of Spades" // go compilers will figure out that card is a string thanks to := operator, only when creating variable
	// card = "Five of diamonds"

	// card := newCard()

	cards := []string{"Ace of diamonds", newCard()}
	cards = append(cards, "Six of Spades")

	for i, card := range cards {
		fmt.Println(i, card)
	}

}

func newCard() string {
	return "Five of Diamonds"
}
