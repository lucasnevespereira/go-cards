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

## Arrays vs Slices

- Array: Fixed length list of things
- Slice: An array that can grow or shrink

Every Element in a <b>slice</b> must be of the same type

## Declare Slices

```
	cards := []string{"Ace of diamonds", newCard()}
```

Appending/pushing new element to a slice

```
	cards = append(cards, "Six of Spades")
```

## Declare Loops

Iterating trough a slice

```
	for i, card := range cards {
		fmt.Println(i, card)
	}
```

<b>i</b> is the index of the element in array. <br>
<b>card</b> is the current card we are iterating over.<br>
<b>range cards</b> takes the slice of cards and loops over it.<br>
<b>fmt.Println()</b> runs this for each card we are iterating.<br>
