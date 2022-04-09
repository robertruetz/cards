package models

import (
	"fmt"
	"github.com/robertruetz/deckOfCards/pkg"
	"strings"
)

type Hand struct {
	DownCard pkg.Card
	UpCards  []pkg.Card
}

func NewHand(down pkg.Card, up pkg.Card) Hand {
	down.FaceUp = false
	up.FaceUp = true
	h := Hand{
		DownCard: down,
		UpCards:  []pkg.Card{up},
	}
	return h
}

func (h *Hand) IsBlackjack() bool {
	if len(h.UpCards) != 1 {
		return false
	}
	if (h.DownCard.IsTenCard() && h.UpCards[0].IsAce()) ||
		(h.UpCards[0].IsTenCard() && h.DownCard.IsAce()){
		return true
	}
	return false
}

// Sum adds the value of each card in the given hand.
// If a card is an Ace, adds 1 or 11 based on whichever gets us closer (but not over) 21.
func (h *Hand) Sum() int {
	sum := 0
	for _, c := range h.UpCards {
		if c.IsAce() {
			sum += handleAce(sum)
			continue
		}
		sum += c.Rank.Value
	}
	sum += h.DownCard.Rank.Value
	return sum
}

func handleAce(sum int) int {
	if sum + 11 > 21 {
		return 1
	}
	return 11
}

func (h *Hand) FlipDownCard() {
	h.DownCard.FaceUp = true
}

func (h *Hand) PrintHand() string {
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("\n\tDown: %s\n", h.DownCard))
	sb.WriteString("\tUp: ")
	for i, c := range h.UpCards {
		sb.WriteString(c.String())
		if i + 1 < len(h.UpCards) {
			sb.WriteString(", ")
		}
	}
	return sb.String()
}
