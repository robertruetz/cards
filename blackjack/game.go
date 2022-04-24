package blackjack

import (
	"fmt"
	"github.com/robertruetz/deckOfCards/blackjack/models"
	"github.com/robertruetz/deckOfCards/blackjack/ui"
	"github.com/robertruetz/deckOfCards/pkg"
)

type Game struct {
	deck *pkg.CardDeck
	players []*Player
	dealer *Player
	ui *ui.Ui
}

func NewGame(ui *ui.Ui, d *pkg.CardDeck, numPlayers int) *Game {
	g := Game{
		ui: ui,
		deck: d,
	}
	plys := make([]*Player, numPlayers)
	for i := 0; i < numPlayers; i++ {
		plys[i] = NewPlayer(models.NewHand(*g.deck.Deal(false), *g.deck.Deal(true)), fmt.Sprintf("Player%v", i + 1))
	}
	g.players = plys
	g.dealer = NewPlayer(models.NewHand(*g.deck.Deal(false), *g.deck.Deal(true)), "dealer")
	return &g
}

func (g *Game) playerCount() int {
	return len(g.players)
}

func (g *Game) PrintHands() {
	logger.Printf("dealer hand: %s\n", g.dealer.Hand.PrintHand())
	g.ui.HorizontalLine()
	for _, p := range g.players{
		logger.Printf("%s hand: %s\n", p.Name, p.Hand.PrintHand())
		g.ui.HorizontalLine()
	}
}
