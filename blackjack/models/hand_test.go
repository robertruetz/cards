package models

import (
	"github.com/robertruetz/deckOfCards/pkg"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	AceOfDiamonds = pkg.NewCard(pkg.DIAMONDS, pkg.Rank{Value: 1, Name: "Ace"})
	KingOfClubs = pkg.NewCard(pkg.CLUBS, pkg.Rank{Value: 10, Name: "King"})
	KingOfDiamonds = pkg.NewCard(pkg.DIAMONDS, pkg.Rank{Value: 10, Name: "King"})
)

func TestHand_IsBlackjack(t *testing.T) {
	h := NewHand(AceOfDiamonds, KingOfClubs)
	assert.True(t, h.IsBlackjack())
}

func TestHand_Sum(t *testing.T) {
	h := NewHand(AceOfDiamonds, KingOfClubs)
	got := h.Sum()
	assert.Equal(t, 21, got)

	h = NewHand(KingOfDiamonds, KingOfClubs)
	got = h.Sum()
	assert.Equal(t, 20, got)
}
