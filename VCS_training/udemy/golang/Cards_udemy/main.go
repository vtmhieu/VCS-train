package main

import "fmt"

// cards project
func main() {
	//var card string = "Ace of Spades"
	//card := "Ace of Spades"
	//only use := when we declare a NEW VARIABLE
	//card = "Five of Diamonds"
	//variable, name card, type string
	//type basic :string , bool, int, float64

	//card := newCard()

	//array and slice
	//array - fixed length array
	//slice = grow or shrink
	type deck []string
	cards := deck{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades")

	//cards.print()

	// for index, card := range cards{}
	//index of the element in array
	// current card
	//range card = take the slice of "cards" and loop over it

	fmt.Println(cards)

}

func newCard() string {
	//need to tell exactly what type of return type
	return "Five of diamonds"
}

// Class object oriented
// Go is not an OOP language so there is no idea of class in Go

// we have this main.go
// next we should have deck.go to describe what a deck is and how it works
// last we shoould have deck_test.go to automatically test the deck
