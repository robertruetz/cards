package blackjack

import (
	"github.com/robertruetz/deckOfCards/blackjack/models"
	"github.com/robertruetz/deckOfCards/blackjack/ui"
	"github.com/robertruetz/deckOfCards/pkg"
)

type Game struct {
	deck *pkg.CardDeck
	player *Player
	dealer *Player
	ui *ui.Ui
}

func NewGame(ui *ui.Ui, d *pkg.CardDeck) *Game {
	g := Game{
		ui: ui,
		deck: d,
	}
	plyr := NewPlayer(models.NewHand(*g.deck.Deal(false), *g.deck.Deal(true)), "player")
	g.player = plyr
	g.dealer = NewPlayer(models.NewHand(*g.deck.Deal(false), *g.deck.Deal(true)), "dealer")
	return &g
}

func (g *Game) PrintHands() {
	logger.Printf("dealer hand: %s\n", g.dealer.Hand.PrintHand())
	g.ui.HorizontalLine()
	logger.Printf("player hand: %s\n", g.player.Hand.PrintHand())
	g.ui.HorizontalLine()
}
