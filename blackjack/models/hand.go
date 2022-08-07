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
	//if len(h.UpCards) != 1 {
	//	return false
	//}
	if (h.DownCard.IsTenCard() && h.UpCards[0].IsAce()) ||
		(h.UpCards[0].IsTenCard() && h.DownCard.IsAce()){
		return true
	}
	return false
}

func (h *Hand) IsBust() bool {
	return h.Sum() > 21
}

// Sum adds the value of each card in the given hand.
// Aces are handled such that a bust is prevented if possible.
func (h *Hand) Sum() int {
	aces := make([]pkg.Card, 0)
	sum := 0
	if h.DownCard.IsAce() {
		aces = append(aces, h.DownCard)
	} else {
		sum += h.DownCard.Rank.Value
	}
	for _, c := range h.UpCards {
		if c.IsAce() {
			aces = append(aces, h.DownCard)
			continue
		}
		sum += c.Rank.Value
	}
	// handle aces
	for _, _ = range aces {
		sum += handleAce(sum)
	}
	return sum
}

// handleAce returns 1 or 11 based on whichever puts the sum closer to (but not over) 21.
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
