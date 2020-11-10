package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
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

type Player struct {
	race     string
	handDeck []Card
}

func (p *Player) addHandDeck(pd *PublicDeck) {
	p.handDeck = append(p.handDeck, pd.drawCard())
}

func (p *Player) handScore() []int {
	var sum int
	var aceCount int
	for _, card := range p.handDeck {
		sum += card.number
		if card.mark == "A" {
			aceCount++
		}
	}
	var handScores []int = []int{sum}
	if aceCount > 0 {
		handScores = append(handScores, (sum - aceCount*10))
	}
	return handScores
}

func (p *Player) maxScore() int {
	maxScore := 0
	handScore := p.handScore()
	for _, s := range handScore {
		if s <= 21 && s > maxScore {
			maxScore = s
		}
	}
	return maxScore
}

func (p *Player) openHandDeck() string {
	var handDeck string
	for _, c := range p.handDeck {
		handDeck = handDeck + " " + c.mark
	}
	return handDeck
}
func (d *PublicDeck) makeDeck() {
	var m string
	var n int
	for s := 1; s <= 4; s++ {
		for i := 1; i <= 13; i++ {
			if i == 1 {
				m = "A"
				n = 11
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

func (d *PublicDeck) drawCard() Card {
	card := d.deck[0]
	d.deck = d.deck[1:]
	return card
}

func main() {
	d := PublicDeck{}
	d.makeDeck()
	d.shuffleDeck()

	p1 := Player{race: "human", handDeck: []Card{}}
	c1 := Player{race: "CPU", handDeck: []Card{}}

	p1.addHandDeck(&d)
	c1.addHandDeck(&d)
	p1.addHandDeck(&d)
	c1.addHandDeck(&d)

	cpuStopFlag := false
	playerStopFlag := false

	for {
		fmt.Print("CPU hand cards: ", c1.handDeck[0].mark)
		fmt.Println(strings.Repeat(" *", (len(c1.handDeck) - 1)))

		fmt.Print("Your hand cards: ")
		for _, c := range p1.handDeck {
			fmt.Print(c.mark, " ")
		}
		fmt.Print("\n")
		fmt.Println("--- --- --- --- ---")
		if c1.maxScore() == 0 {
			cpuStopFlag = true
		} else if c1.maxScore() < 17 {
			c1.addHandDeck(&d)
			fmt.Println("CPU is HIT")
		} else {
			cpuStopFlag = true
			fmt.Println("CPU is STAND")
		}
		if p1.maxScore() == 0 {
			fmt.Println("Your card is Burst!")
			playerStopFlag == true
		}
		if playerStopFlag == false {
			fmt.Print("Are you HIT ? (y/n)")
			var yn string
			fmt.Scan(&yn)
			if yn == "y" {
				p1.addHandDeck(&d)
			} else if yn == "n" {
				playerStopFlag = true
			}
		}
		if cpuStopFlag && playerStopFlag {
			break
		}
	}
	fmt.Println("Open!")
	time.Sleep(2 * time.Second)
	fmt.Println("=== === === === ===", "\n")
	//勝利判定
	if p1.maxScore() == 0 {
		fmt.Println("You lose. Your card is Burst.")
	} else if c1.maxScore() >= p1.maxScore() {
		fmt.Println("You lose")
	} else {
		fmt.Println("You win!")
	}
	fmt.Println("CPU score: ", c1.maxScore(), ",CPU hand cards: ", c1.openHandDeck())
	fmt.Println("Your score: ", p1.maxScore(), ",Your hand cards: ", p1.openHandDeck())
}
