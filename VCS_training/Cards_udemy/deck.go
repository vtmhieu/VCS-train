package main

import "fmt"

//ioutil package implemaent some I/O utility fucntions
//create a new type of 'deck'
//which is a slice of strings
// fucntions WriteFile writes data to a file named by filename.
//If teh file does not exist, WriteFile creates it with permissions perm; otherwise WriteFile truncates it before writing
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"A", "B"}
	cardValues := []string{"Hearts", "Diamonds"}
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" "+suit)
		}
	}

	return cards
	//has to return
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
func deal(d deck, handSize int) (deck, deck) {
	//(deck,deck) is beleived to tell go we going to return 2 deck
	return d[:handSize], d[handSize:]
}

func main() {
	//def new deck of cards
	cards := newDeck()
	//print out all the card in deck
	cards.print()
	//print out cards 0 and 1
	cards[0:2].print()
	//print out all the card in deck from 1st
	cards[1:].print()

	//print out deal func
	hand, remaindingCards := deal(cards, 1)

	hand.print()
	remaindingCards.print()
}

//
//21
