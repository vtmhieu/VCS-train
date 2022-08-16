package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// ioutil package implemaent some I/O utility fucntions
// create a new type of 'deck'
// which is a slice of strings
// fucntions WriteFile writes data to a file named by filename.
// If teh file does not exist, WriteFile creates it with permissions perm; otherwise WriteFile truncates it before writing
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

// func WriteFile()

//data []byte is slice of bytes -> in order to write file we need to
// change the deck of card type into slice of bytes
//"Hi there!" string -> [72 105 32 116 104 101 114 101 33] byte slice
// asciitable.com can be used as Decimal

//[]byte("Hi there!") convert string into byte

// perm os.FileMode is permissions on ...
// as who can access to the file, who can read it

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// change  deck into string

func (d deck) savetoFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
	//0666 means anyone can read it
}
func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	//bs = byteslice, err= error
	//err, if nothing went wrong it will hace a value of 'nil' = no value
	if err != nil {
		//Option 1: log the error and return to a call to newDeck()

		// return to newDeck()
		//Option 2: Log the error and entirely quit the program
		fmt.Println("error:", err)
		os.Exit(1) //log out and quit the program
	}

	ss := strings.Split(string(bs), ",")
	return deck(ss)

}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
func main() {
	//def new deck of cards
	cards := newDeck()
	// //print out all the card in deck
	// cards.print()
	// //print out cards 0 and 1
	// cards[0:2].print()
	// //print out all the card in deck from 1st
	// cards[1:].print()

	// //print out deal func
	// hand, remaindingCards := deal(cards, 1)

	// hand.print()
	// remaindingCards.print()
	// greeting := "Hi there!"
	// fmt.Println([]byte(greeting))
	// fmt.Println(cards.toString())
	// cards.savetoFile("my_cards")
	// cards := newDeckFromFile("my_cards")
	cards.shuffle()
	cards.print()
}

//
//21
