/*
Create type "deck" that extends a slice of strings
That also contains functions exclusive to type deck
*/

// The main package that allows deck.go to be executed within main.go
package main

// Various imports that allows sandard go library functions to be used in this file
import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// initialises a new type called deck that extends a slice of strings
type deck []string

// creates a new deck of playing cards - must return type deck
func newDeck() deck {
	// initialise variable cards as type deck
	cards := deck{}
	// create a variable as a string slice with values of playing card suits
	cardSuits := []string{"Hearts", "Diamonds", "Spades", "Clubs"}
	// create a variable as a string slice with values of playing card values
	cardValues := []string{"Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King", "Ace"}

	// for every suit in the variable card suits run the nested routine
	for _, suit := range cardSuits {
		// for every value in the variable card values run the nest routine
		for _, value := range cardValues {
			// variable cards (type deck) add current value, and then current suit to cards variable
			cards = append(cards, value+" of "+suit)
		}
	}

	// return cards (type deck)
	return cards
}

// function that prints all values of slice strings
func (d deck) print() {
	// for each key/index (i) and card (value) in d (type deck) (d referencing type deck variable passed into print function)
	for i, card := range d {
		// print index (i) and value (card) to console
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize] /* Hand */, d[handSize:] /* Deck after deal */
}

func (d deck) toString() string {
	return strings.Join(d, ",")
}

func (d deck) writeToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

func readFile(fileName string) deck {
	bs, err := ioutil.ReadFile(fileName)

	if err != nil {
		fmt.Println("There was an issue regarding reading filename " + fileName)
		fmt.Println("Error Message:", err)

		os.Exit(1)
	}

	if bs == nil {
		fmt.Println("File read is empty, creating new deck...")
		return newDeck()
	}

	s := strings.Split(string(bs), ",")

	return deck(s)
}

func (d deck) shuffle() deck {

	source := rand.NewSource(time.Now().UnixNano()) // create new source with the value of now in nanoseconds
	r := rand.New(source)                           // add source to rand struct

	for i := range d {
		position := r.Intn(len(d) - 1)
		d[i], d[position] = d[position], d[i] // swap positions of slice index (note both must be swapped or you copy the former slice value)
	}

	return d
}
