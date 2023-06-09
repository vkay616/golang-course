package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
)

// custom type deck
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Clubs", "Hearts"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	dealtCards := deck{}
	remainingCards := d

	for i := 1; i <= handSize; i++ {

		n := rand.Intn(len(d) - 1)

		dealtCards = append(dealtCards, remainingCards[n])
		remainingCards.removeElement(n)

	}

	return dealtCards, remainingCards
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func loadDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("ERROR:", err)
		os.Exit(1)
	}

	return deck(strings.Split(string(bs), ","))
}

func (d deck) shuffle() {
	for i := range d {
		n := rand.Intn(len(d) - 1)
		d[i], d[n] = d[n], d[i]
	}
}

func (d deck) removeElement(index int) deck {
	d[index] = d[len(d)-1]
	return d[:len(d)-1]
}

func (d deck) containsDuplicates() bool {
	checkedCards := deck{}

	for _, card := range d {
		for _, checkedCard := range checkedCards {
			if card == checkedCard {
				return true
			}
		}
		checkedCards = append(checkedCards, card)
	}

	return false
}
