package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of Deck
// Which is a slice of strings
type deck []string

// Function to print all

func (d deck) print() {
	for _, j := range d {
		fmt.Println(j)
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
	cardSuit := []string{"♠", "♥", "♦", "♣"}
	cardValue := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "💂", "👸", "🦁"}
	cards := deck{}

	for _, suit := range cardSuit {
		for _, value := range cardValue {
			cards = append(cards, value+" of "+suit)
		}

	}
	return cards
}

//Function to return hand
func hand(d deck, handSize int) (deck, deck) {
	return d[0:handSize], d[handSize:len(d)]

}

// Function to convert the deck to single string
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// Function to write the data into a file
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)

}

// Function to read from File
func newDeckFromFile(fileName string) deck {
	bs, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	allcards := string(bs)                   // coverting byte slice to string
	strslice := strings.Split(allcards, ",") //spliting the string by ',' and converting it into slice of string
	return deck(strslice)
}
