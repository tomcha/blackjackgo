package main

import (
	"strconv"
)

type Card struct {
	suite  int
	number int
	mark   string
}

type PublicDeck struct {
	deck []Card
}

func (d *PublicDeck) makeDeck() {
	var m string
	var n int
	for s := 1; s <= 4; s++ {
		for i := 1; i <= 13; i++ {
			if i == 1 {
				m = "A"
				n = i
			} else if i == 11 {
				m = "J"
				n - 10
			} else if i == 12 {
				m = "Q"
				n = 10
			} else if i == 13 {
				m = "K"
				n = 10
			} else {
				m = strconv.Itoa(i)
				n = i
			}
			d.deck = append(d.deck, Card{suit})
		}
	}
}

func (d *PublicDeck) shuffleDeck() {
}

func main() {
}
