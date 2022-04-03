package pkg

import (
	"math/rand"
	"time"
)

const TotalCards = 52

type Suit string

const (
	DIAMONDS Suit = "DIAMONDS"
	CLUBS = "CLUBS"
	SPADES = "SPADES"
	HEARTS = "HEARTS"
)

type CardDeck struct {
	cards []Card
	currentCard int
}

func NewDeckOfCards() CardDeck {
	d := CardDeck{
		cards: make([]Card, TotalCards),
	}
	suits := []Suit{DIAMONDS, CLUBS, SPADES, HEARTS}
	dIx := 0
	for _, s := range suits {
		for i := 1; i <= 13; i++ {
			d.cards[dIx] = NewCard(s, i)
			dIx++
		}
	}
	return d
}

func (d *CardDeck) swap(pos1 int, pos2 int) {
	hold := d.cards[pos1]
	d.cards[pos1] = d.cards[pos2]
	d.cards[pos2] = hold
}

// Shuffle swaps the card at each position in the deck with a random card from the remaining positions.
func (d *CardDeck) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < TotalCards; i++ {
		rando := rand.Intn(TotalCards- i) + i
		d.swap(i, rando)
	}
}

// Deal returns the card at the index currentCard for the given deck.
// Nil is returned if the currentCard is greater than TotalCards signifying an empty deck
func (d *CardDeck) Deal() *Card {
	var c *Card
	if d.currentCard < TotalCards {
		c = &d.cards[d.currentCard]
		d.currentCard++
	}
	return c
}

type Card struct {
	Suit Suit
	Rank int
}

func NewCard(suit Suit, rank int) Card {
	return Card{
		Suit: suit,
		Rank: rank,
	}
}

// TODO: Implement a ShoeOfCards -- n number of decks all shuffled together that can be reshuffled at a given interval
