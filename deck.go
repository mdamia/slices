package main

import (
	"log"
	"math/rand"
)

type deck []string

// newDeck
func newDeck() deck {

	cards := deck{}

	cardSuits := []string{"Spades", "Diamons", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Tow", "Three", "Four"}
	for _, cs := range cardSuits {
		for _, cv := range cardValues {
			cards = append(cards, cv+" of "+cs)

		}
	}

	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func shuffleDeck(d deck) []string {
	dest := make([]string, len(d))
	perm := rand.Perm(len(d))
	for i, p := range perm {
		dest[p] = d[i]
	}
	return dest
}

func (d deck) print() {
	for _, card := range d {
		log.Println(card)
	}
}
