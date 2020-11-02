package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
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
				n = 10
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
			d.deck = append(d.deck, Card{suite: s, number: n, mark: m})
		}
	}
}

func (d *PublicDeck) shuffleDeck() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d.deck), func(i, j int) { d.deck[i], d.deck[j] = d.deck[j], d.deck[i] })
}

func main() {
	d := PublicDeck{}
	d.makeDeck()
}
