package blackjack

import (
	"fmt"
	cards "github.com/robertruetz/deckOfCards/pkg"
	"strings"
)

type Hand struct {
	downCard cards.Card
	upCards []cards.Card
}

func NewHand(down cards.Card, up cards.Card) Hand {
	down.FaceUp = false
	up.FaceUp = true
	h := Hand{
		downCard: down,
		upCards: []cards.Card{ up },
	}
	return h
}

type Player struct {
	hand Hand
}

func (p *Player) PrintHand() string {
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("\n\tDown: %s\n", p.hand.downCard))
	sb.WriteString("\tUp: ")
	for i, c := range p.hand.upCards {
		sb.WriteString(c.String())
		if i + 1 < len(p.hand.upCards) {
			sb.WriteString(", ")
		}
	}
	return sb.String()
}

func NewPlayer(hand Hand) *Player {
	p := Player{
		hand: hand,
	}
	return &p
}

func (p *Player) DealCard(card cards.Card) {
	if p.hand.upCards == nil {
		p.hand.upCards = make([]cards.Card, 0)
	}
	p.hand.upCards = append(p.hand.upCards, card)
}

func Play(){
	deck := cards.NewCardDeck()
	deck.Shuffle()
	p1 := NewPlayer(NewHand(*deck.Deal(false), *deck.Deal(true)))
	dealer := NewPlayer(NewHand(*deck.Deal(false), *deck.Deal(true)))

	p1.DealCard(*deck.Deal(true))

	fmt.Printf("Player 1 hand: %s\nDealer hand: %s\n", p1.PrintHand(), dealer.PrintHand())
	// TODO:
		// decision p1
		// decision dealer
		// resolve game
}

/* Dealer Rules
When the dealer has served every player, the dealers face-down card is turned up.
If the total is 17 or more, it must stand.
If the total is 16 or under, they must take a card.
The dealer must continue to take cards until the total is 17 or more, at which point the dealer must stand.
If the dealer has an ace, and counting it as 11 would bring the total to 17 or more (but not over 21),the dealer must count the ace as 11 and stand.
 */
