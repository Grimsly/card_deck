package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

// Create a new type of deck
// Which is a slice of strings
type deck []string

func newDeck() deck{
	// := initializes a variable and also assigns it a value
	// Based on the value assigned, it can guess what type the variable is supposed to be (similar to Typescript?)
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value + " of " + suit)
		}
 	}

	return cards
}

/* 
Putting a field before the function name turns the function into a receiver function.
This allows only variables of deck type to use the function.
Similar to methods in OOP classes
 */
func (d deck) print(){ 
	for i, card := range d{
		fmt.Println(i, card)
	}
}

// Convention for Golang is to have the first letter of the type be the variable name
func (d deck) deal(handSize int) (deck, deck){
	// Equivalent to Javascript's Array.slice() function
	// Returns the items in the list from the start to the given handSize
	// and the the items starting from the index of the handSize to the end
	// This does not manipulate the original slice
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	// Concatenates the deck (slice) into a string, with its items separated by a comma
	return strings.Join(d, ",")
}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	// Not necessary to convert slice into deck type since Go can infer that deck is a slice
	return deck(s)
}

func (d deck) shuffle() {
	// As of 2022, creating a random number generator on program startup is unnecessary as the seed is now not a constant
	// source := rand.NewSource(time.Now().UnixNano())
	// r := rand.New(source)

	for i := range d {
		// len() is the same as Array.length
		newPosition := rand.Intn(len(d) - 1)
		// newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}