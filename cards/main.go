package main

import "fmt"

func main() {
	// var card string = "Ace of Spades"
	// card := newCard()
	// fmt.Println(card)
	// cards := deck{newCard(), "Ace of Diamonds", newCard()}
	// cards = append(cards, "Five of Club")
	cards := newDeck()

	// fmt.Println(cards.toString())

	// cards.saveToFile("saved_cards")

	// cards := loadDeckFromFile("saved_cards")

	// cards.shuffle()

	// cards.print()

	hand, remainingCards := deal(cards, 10)

	fmt.Println("All Cards:")
	cards.print()

	fmt.Println("Cards in Hand after dealing:")
	hand.print()

	fmt.Println("Remaining cards in pile:")
	remainingCards.print()

	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }

	// fmt.Println(cards)
}
