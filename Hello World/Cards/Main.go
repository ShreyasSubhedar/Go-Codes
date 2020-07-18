package main

import "fmt"

func main() {
	// Creating a deck
	cards := newDeck()
	// Printing out the deck
	cards.print()
	// Shuffle a deck
	cards.shuffle()

	fmt.Println("Shuffling the all cards")
	// Printing out the deck
	cards.print()
	fmt.Println("Creating new Card Deck")

}
