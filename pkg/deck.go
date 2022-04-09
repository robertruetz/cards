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
			var rank Rank = GetRank(i)
			d.cards[dIx] = NewCard(s, rank)
			dIx++
		}
	}
	return d
}

func (d *CardDeck) DeepCopy() CardDeck {
	if d.cards == nil {
		return CardDeck{cards: []Card(nil)}
	}
	cd := CardDeck{
		cards: make([]Card, len(d.cards)),
		currentCard: d.currentCard,
	}
	for ix, c := range d.cards {
		cd.cards[ix] = NewCard(c.Suit, c.Rank)
		cd.cards[ix].FaceUp = c.FaceUp
	}
	return cd
}

// swap replaces the value at pos1 with the value at pos2 and vice versa.
func (d *CardDeck) swap(pos1 int, pos2 int) {
	hold := d.cards[pos1]
	d.cards[pos1] = d.cards[pos2]
	d.cards[pos2] = hold
}

// Shuffle swaps the card at each position in the deck with a random card from the remaining possible positions.
// Simulates shuffling a deck of cards
func (d *CardDeck) Shuffle() error {
	if d.cards == nil || len(d.cards) == 0 {
		return fmt.Errorf("card deck was empty")
	}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < TotalCards; i++ {
		rando := rand.Intn(TotalCards- i) + i
		d.swap(i, rando)
	}
	return nil
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
	Rank Rank
	FaceUp bool
}

func (c Card) String() string {
	return fmt.Sprintf("%s of %c", c.Rank.String(), symbols[string(c.Suit)])
}

// IsTenCard returns true if the card is 10, Jack, Queen, or King.
func (c *Card) IsTenCard() bool {
	return c.Rank.Value == 10
}

func (c *Card) IsAce() bool {
	return c.Rank.Name == "Ace"
}

func NewCard(suit Suit, rank Rank) Card {
	return Card{
		Suit: suit,
		Rank: rank,
		FaceUp: false,
	}
}

type Rank struct {
	Name string
	Value int
	AltValue int
}

func GetRank(value int) Rank{
	r := Rank{
		Name: rankSwap(value),
		Value: value,
	}
	if value >= 10 {
		r.Value = 10
	}
	if value == 1 {
		r.AltValue = 11
	}
	return r
}

func (r *Rank) String() string {
	return r.Name
}

// rankSwap swaps out the Rank of the given card for the appropriate string value.
// Handles Ace, Jack, Queen, King
func rankSwap(value int) string {
	switch value {
	case 1:
		return "Ace"
	case 11:
		return "Jack"
	case 12:
		return "Queen"
	case 13:
		return "King"
	default:
		return strconv.Itoa(value)
	}
}

// TODO: Implement a ShoeOfCards -- n number of decks all shuffled together that can be reshuffled at a given interval
