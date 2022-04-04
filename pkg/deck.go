package pkg

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

const TotalCards = 52

type Suit string

const (
	DIAMONDS Suit = "DIAMONDS"
	CLUBS Suit = "CLUBS"
	SPADES Suit = "SPADES"
	HEARTS Suit = "HEARTS"
)

var symbols map[string]rune = map[string]rune{
	"SPADES": '♠',
	"DIAMONDS": '♦',
	"HEARTS": '♥',
	"CLUBS": '♣',
}

type CardDeck struct {
	cards []Card
	currentCard int
}

// NewCardDeck returns an ordered deck of playing cards.
func NewCardDeck() CardDeck {
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

// swap replaces the value at pos1 with the value at pos2 and vice versa.
func (d *CardDeck) swap(pos1 int, pos2 int) {
	hold := d.cards[pos1]
	d.cards[pos1] = d.cards[pos2]
	d.cards[pos2] = hold
}

// Shuffle swaps the card at each position in the deck with a random card from the remaining possible positions.
// Simulates shuffling a deck of cards
func (d *CardDeck) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < TotalCards; i++ {
		rando := rand.Intn(TotalCards- i) + i
		d.swap(i, rando)
	}
}

// Deal returns the card at the index currentCard for the given deck.
// Nil is returned if the currentCard is greater than TotalCards signifying an empty deck
func (d *CardDeck) Deal(faceUp bool) *Card {
	var c *Card
	if d.currentCard < TotalCards {
		c = &d.cards[d.currentCard]
		c.FaceUp = faceUp
		d.currentCard++
	}
	return c
}

type Card struct {
	Suit Suit
	Rank int
	FaceUp bool
}

func (c Card) String() string {
	return fmt.Sprintf("%v of %c", c.rankSwap(), symbols[string(c.Suit)])
}

func NewCard(suit Suit, rank int) Card {
	return Card{
		Suit: suit,
		Rank: rank,
		FaceUp: false,
	}
}

// rankSwap swaps out the Rank of the given card for the appropriate string value.
// Handles Ace, Jack, Queen, King
func (c Card) rankSwap() string {
	switch c.Rank {
	case 1:
		return "Ace"
	case 11:
		return "Jack"
	case 12:
		return "Queen"
	case 13:
		return "King"
	default:
		return strconv.Itoa(c.Rank)
	}
}

// TODO: Implement a ShoeOfCards -- n number of decks all shuffled together that can be reshuffled at a given interval
