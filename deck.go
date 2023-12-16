package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type deck []string //create a new type of 'deck, which is a slice of string

//Receiver function, d is like this or self in other languages. You can give it any name, even say 'this'.
//But as per recomendation, this should be the first one/two letters of the custom type. 
func (d deck) print() { 
	for i, card := range d { 
		fmt.Println(i, card)
	}
}

func newDeck() deck {
	cards := deck{} //you can initialise with an empty object of custom type
	cardSuits := []string{"Spade", "Diamond"}
	cardValues := []string{"Ace", "Two", "Queen"}
	for _, suit := range cardSuits { //use _ in for loop, if you dont want to use index, else compiler will complain about unused var
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards;
}

// Convert deck to a array of string (this is possible as deck is ultimately an array of string only), 
// and then to a comma seperated string value
func (d deck) toString() string { 
	return strings.Join([]string(d), ",")
}

//Convert string to byte[], as that is what writeFile method of ioutil package expects.
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

//Read the deck saved in file, with error handling
func (d deck) newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1) //Exit the program
	}

	s := strings.Split(string(bs), ",")
	return deck(s)
}