// Reminder: Package main defines an executable output
package main

func main() {

	//cards := newDeck()

	// // declareing 2 variables, cause we have 2 return values
	//hand, remainingCards := deal(cards, 5)

	// // They are type deck so they have access to print() that is a receiver func
	// hand.print()
	// remainingCards.print()

	// cards.saveToFile("my_cards.txt")

	cards2 := newDeckFromFile("my_cards.txt")
	cards2.shuffle()
	cards2.print()

}
