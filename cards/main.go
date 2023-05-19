package main

func main() {
	// var card string = "Ace of Spades"
	// card := newCard()
	// fmt.Println(card)
	cards := deck{newCard(), "Ace of Diamonds", newCard()}
	cards = append(cards, "Five of Club")

	cards.print()

	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }

	// fmt.Println(cards)
}

func newCard() string {
	return "Ace of Hearts"
}
