package deckOfCards

import (
	"math/rand"
	"time"
)

const TotalCards = 52

type Suit int

const (
	DIAMONDS Suit = iota
	CLUBS
	SPADES
	HEARTS
)

type DeckOfCards struct {
	cards []Card
	currentCard int
}

func NewDeckOfCards() DeckOfCards {
	d := DeckOfCards{
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

func (d *DeckOfCards) swap(pos1 int, pos2 int) {
	hold := d.cards[pos1]
	d.cards[pos1] = d.cards[pos2]
	d.cards[pos2] = hold
}

// Shuffle swaps the card at each position in the deck with a random card from the remaining positions.
func (d *DeckOfCards) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < TotalCards; i++ {
		rando := rand.Intn(TotalCards - i) + i
		d.swap(i, rando)
	}
}

// Deal returns the card at the index currentCard for the given deck.
// Nil is returned if the currentCard is greater than TotalCards signifying an empty deck
func (d *DeckOfCards) Deal() *Card {
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
