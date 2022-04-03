package blackjack

import (
	"errors"
	"fmt"
	cards "github.com/robertruetz/deckOfCards/pkg"
	"github.com/robertruetz/deckOfCards/util"
)

const InitHandSize = 2

type Player struct {
	hand []cards.Card
}

func NewPlayer() *Player {
	p := Player{
		hand: make([]cards.Card, 0, InitHandSize),
	}
	return &p
}

func (p *Player) DealCard(card cards.Card) error {
	if len(p.hand) < InitHandSize {
		p.hand = append(p.hand, card)
		return nil
	}
	return errors.New("hand is full")
}

func Play(){
	deck := cards.NewDeckOfCards()
	deck.Shuffle()
	p1 := NewPlayer()
	p2 := NewPlayer()
	for i := 0; i < InitHandSize; i++ {
		err := p1.DealCard(*deck.Deal())
		if err != nil {
			panic(err.Error())
		}
		err = p2.DealCard(*deck.Deal())
		if err != nil {
			panic(err.Error())
		}
	}
	fmt.Printf("Player 1 hand: %s\n Player 2 hand: %s\n", util.JSONify(p1.hand), util.JSONify(p2.hand))
	// TODO:
		// decision p1
		// decision p2
		// resolve game
}
