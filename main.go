// Reminder: Package main defines an executable output
package main

func main() {

	cards := deck{"Ace of diamonds", newCard()}
	cards = append(cards, "Six of Spades")

	cards.print()

}

func newCard() string {
	return "Five of Diamonds"
}
