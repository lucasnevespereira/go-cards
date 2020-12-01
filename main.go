// Reminder: Package main defines an executable output
package main

import "fmt"

func main() {

	cards := newDeck()

	// // declareing 2 variables, cause we have 2 return values
	//hand, remainingCards := deal(cards, 5)

	// // They are type deck so they have access to print() that is a receiver func
	// hand.print()
	// remainingCards.print()

	fmt.Println(cards.toString())
}
