package blackjack

import (
	cards "github.com/robertruetz/deckOfCards/pkg"
	"log"
)

type IStrategy interface {
	Employ() int
}

/*
https://bicyclecards.com/how-to-play/blackjack/

Dealer Rules
	When the dealer has served every player, the dealers face-down card is turned up.
	If the total is 17 or more, it must stand.
	If the total is 16 or under, they must take a card.
	The dealer must continue to take cards until the total is 17 or more, at which point the dealer must stand.
	If the dealer has an ace, and counting it as 11 would bring the total to 17 or more (but not over 21),the dealer must count the ace as 11 and stand.
*/

type DealerStrategy struct {
	deck *cards.CardDeck
	player *Player
	logger *log.Logger
}

func NewDealerStrategy(deck *cards.CardDeck, player *Player, logger *log.Logger) DealerStrategy {
	ds := DealerStrategy{
		deck: deck,
		player: player,
		logger: logger,
	}
	return ds
}

// Employ employs the official Dealer Rules from https://bicyclecards.com/how-to-play/blackjack/
func (d *DealerStrategy) Employ() int {
	sum := 0
	hand := &d.player.Hand
	stand := false
	for ; !stand; {
		sum := hand.Sum()
		logger.Printf("%s has %v\n", d.player.Name, sum)
		if sum > 21 {
			logger.Printf("%s busts\n", d.player.Name)
			stand = true
		} else if sum < 17 {
			nCard := *d.deck.Deal(true)
			logger.Printf("%s hits: %s\n", d.player.Name, nCard)
			d.player.DealCard(nCard)
		} else if sum >= 17 {
			logger.Printf("%s stands\n", d.player.Name)
			stand = true
		}
	}
	return sum
}
