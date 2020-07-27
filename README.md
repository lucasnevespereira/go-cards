# go-cards

## Declare Variables

```
var card string = "Ace of Spades"

```

or

Go compilers will figure out that card is a string thanks to := operator (only when creating variable)

```
card := "Ace of Spades"
card = "Five of diamonds"
```

Also we can declare it before assign it

```
  var deckSize int
  deckSize = 52
```

## Declare Functions

```
func newCard() string {
	return "Five of Diamonds"
}
```
