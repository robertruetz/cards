package pkg

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewShoe(t *testing.T) {
	s := NewShoe(1)
	assert.Equal(t, StandardDeckSize, s.Deck.Length())

	got := NewShoe(6)
	assert.Equal(t, 6 * StandardDeckSize, got.Deck.Length())
}