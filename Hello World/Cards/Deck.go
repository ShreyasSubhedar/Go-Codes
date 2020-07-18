package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Create a new type of Deck
// Which is a slice of strings
type deck []string

// Function to print all

func (d deck) print() {
	for i, j := range d {
		fmt.Println(i, j)
	}
}

// Function to shuffle the deck
func (d deck) shuffle() {
	rand.Seed(time.Now().Unix())

	rand.Shuffle(len(d), func(i, j int) {
		d[i], d[j] = d[j], d[i]
	})
}

// Function should return a new deck
func newDeck() deck {
	cardSuit := []string{"Spades", "Hearts", "Diamonds", "Club"}
	cardValue := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	cards := deck{}

	for _, suit := range cardSuit {
		for _, value := range cardValue {
			cards = append(cards, value+" of "+suit)
		}

	}

	return cards
}
