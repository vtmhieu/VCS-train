package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", lend(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "Ace of Spades" {
		t.Errorf("Expected ace of Spades, but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckTestFromFile(t *testing.T) {

}
func main() {}
