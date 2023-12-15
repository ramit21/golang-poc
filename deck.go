package main

import "fmt"

type deck []string //create a new type of 'deck, which is a slice of string

//Receiver function, d is like this or self in other languages. You can give it any name, even say 'this'.
//But as per recomendation, this should be the first one/two letters of the custom type. 
func (d deck) print() { 
	for i, card := range d { 
		fmt.Println(i, card)
	}
}

func newDeck() deck{
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