package main

import "fmt"

func main() {
	// var card string = "Ace of Spades"
	// card := newCard()
	// fmt.Println(card)
	cards := []string{newCard(), "Ace of Diamonds", newCard()}
	cards = append(cards, "Five of Club")

	for i, card := range cards {
		fmt.Println(i, card)
	}

	// fmt.Println(cards)
}

func newCard() string {
	return "Ace of Hearts"
}
