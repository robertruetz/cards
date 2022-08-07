package blackjack

import (
	"fmt"
	"github.com/robertruetz/deckOfCards/blackjack/models"
	"github.com/robertruetz/deckOfCards/blackjack/ui"
	cards "github.com/robertruetz/deckOfCards/pkg"
	"log"
	"os"
	"strings"
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
	gui := ui.NewUi(os.Stdout)
	gui.PrintBanner()
	shoe := cards.NewShoe(6)
	sErr := shoe.Shuffle()
	if sErr != nil {
		logger.Printf("error shuffling deck: %s", sErr.Error())
	}
	game := NewGame(&gui, &shoe.Deck)

	dealerHasBlackJack := false
	if game.dealer.Hand.IsBlackjack() {
		dealerHasBlackJack = true
		game.dealer.Hand.FlipDownCard()
	}

	game.PrintHands()
	stands := false
	for !stands {
		sum := game.player.Hand.Sum()
		logger.Printf("%s has %v\n", game.player.Name, sum)
		s, sErr := gui.ReadLn("Would you like to hit (h) or stand (s): ")
		if sErr != nil {
			panic(sErr)
		}
		switch strings.ToLower(s) {
		case "hit":
		case "h":
			nCard := game.deck.Deal(true)
			game.player.DealCard(*nCard)
			gui.Writef("%s hits: \n\tdealt: %s\n", game.player.Name, nCard)
		case "stand":
		case "s":
			stands = true
		default:
			gui.Writef("%s is not a valid command. Try again.\n", s)
			continue
		}
		// check for bust
		if game.player.Hand.IsBust() {
			gui.Writef("%s busts\n", game.player.Name)
			stands = true
		}
	}

	game.ui.HorizontalLine()
	if !dealerHasBlackJack {
		// handle dealer using dealer strategy
		d := NewDealerStrategy(game.deck, game.dealer, logger)
		d.Employ()
	}
	game.ui.HorizontalLine()
	game.ui.Writef("dealer has:\t%v\nplayer has:\t%v\n", game.dealer.Hand.Sum(), game.player.Hand.Sum())
}

func printHands(player models.Hand, dealer models.Hand) {
	fmt.Printf("Player 1 Hand: %s\nDealer Hand: %s\n", player.PrintHand(), dealer.PrintHand())
}