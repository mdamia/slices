package main

import (
	"log"
)

func main() {
	cards := newDeck()

	log.Println(writeByte(cards))

	// // shuffl deck
	// perm := shuffleDeck(cards)
	// // return hand and remainign cards in deck
	// hand, deck := deal(perm, 1)

}

func writeByte(d deck) (b []byte) {

	for _, card := range d {
		b = []byte(card)
	}

	return b
}
