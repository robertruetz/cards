package blackjack

import (
	"fmt"
	"github.com/robertruetz/deckOfCards/blackjack/models"
	"github.com/robertruetz/deckOfCards/blackjack/ui"
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
	gui := ui.NewUi(os.Stdout)
	gui.PrintBanner()
	pCnt, pCntErr := gui.GetPlayerCount()
	if pCntErr != nil {
		panic (pCntErr)
	} else if pCnt == 0 {
		panic("player count was zero")
	}
	shoe := cards.NewShoe(6)
	sErr := shoe.Shuffle()
	if sErr != nil {
		logger.Printf("error shuffling deck: %s", sErr.Error())
	}
	game := NewGame(&gui, &shoe.Deck, pCnt)

	dealerHasBlackJack := false
	if game.dealer.Hand.IsBlackjack() {
		dealerHasBlackJack = true
		game.dealer.Hand.FlipDownCard()
	}

	game.PrintHands()
	// handle player using dealer strategy
	for _, p := range game.players {
		s := NewDealerStrategy(game.deck, p, logger)
		s.Employ()
		game.ui.HorizontalLine()
	}
	if !dealerHasBlackJack {
		// handle dealer using dealer strategy
		d := NewDealerStrategy(game.deck, game.dealer, logger)
		d.Employ()
	}
}

func printHands(player models.Hand, dealer models.Hand) {
	fmt.Printf("Player 1 Hand: %s\nDealer Hand: %s\n", player.PrintHand(), dealer.PrintHand())
}