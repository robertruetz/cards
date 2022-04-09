package blackjack

import (
	"fmt"
	"github.com/robertruetz/deckOfCards/blackjack/models"
	cards "github.com/robertruetz/deckOfCards/pkg"
	"log"
	"os"
)

var logger = log.New(os.Stdout, "", 0)

type Player struct {
	Hand models.Hand
	Name string
}

func NewPlayer(hand models.Hand, name string) *Player {
	p := Player{
		Hand: hand,
		Name: name,
	}
	return &p
}

func (p *Player) DealCard(card cards.Card) {
	if p.Hand.UpCards == nil {
		p.Hand.UpCards = make([]cards.Card, 0)
	}
	p.Hand.UpCards = append(p.Hand.UpCards, card)
}

func Play(){
	deck := cards.NewCardDeck()
	deck.Shuffle()
	p1 := NewPlayer(models.NewHand(*deck.Deal(false), *deck.Deal(true)), "player")
	dealer := NewPlayer(models.NewHand(*deck.Deal(false), *deck.Deal(true)), "dealer")
    printHands(p1.Hand, dealer.Hand)
	dealerHasBlackJack := false
	if dealer.Hand.IsBlackjack() {
		dealerHasBlackJack = true
		dealer.Hand.FlipDownCard()
	}
	// handle player using dealer strategy
	s := NewDealerStrategy(&deck, p1, logger)
	s.Employ()
	if !dealerHasBlackJack {
		// handle dealer using dealer strategy
		d := NewDealerStrategy(&deck, dealer, logger)
		d.Employ()
	}
}

func printHands(player models.Hand, dealer models.Hand) {
	fmt.Printf("Player 1 Hand: %s\nDealer Hand: %s\n", player.PrintHand(), dealer.PrintHand())
}