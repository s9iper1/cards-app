package main

import (
	"fmt"
	"strings"
)

// create a new type of deck which is a slice of string

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuites := []string{"Spades", "Diamonds", "Heats", "clubs"}
	cardSuiteValue := []string{"ace", "Two", "Three", "Four"}

	for _, i := range cardSuites {
		for _, v := range cardSuiteValue {
			cards = append(cards, v+" of "+i)
		}

	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) tostring() string {
	return strings.Join([]string(d), ",")
}

// func (d deck) printMyMessage() {
// 	fmt.Println("This is a am testing the receiver function")
// }
