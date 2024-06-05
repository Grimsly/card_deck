package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	// No strict equality (===) for Go as they are all equality operators are treated strictly
	// There are no truey/falsey values
	// nil is not the same as false in this language
	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be Ace of Spades, but got %v", d[0])
	}

	if d[len(d) - 1] != "Four of Clubs" {
		t.Errorf("Expected last card to be Four of Clubs, but got %v", d[len(d) - 1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	// Remove the _decktesting file at the start of the test in case a previous test didn't delete it
	os.Remove("_decktesting")

	d := newDeck()
	d.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if (len(loadedDeck) != 16){
		t.Errorf("Expected 16 cards in deck, got %v", len(loadedDeck))
	}

	// Remove the _decktesting file at the end of the test
	os.Remove("_decktesting")
}