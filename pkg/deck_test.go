package pkg

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCardDeck_Shuffle(t *testing.T) {
	type fields struct {
		cardDeck CardDeck
	}
	tests := []struct {
		name   string
		fields fields
		expectError bool
	}{
		{ name: "happy path", fields: fields{cardDeck: NewCardDeck()}},
		{ name: "card deck empty", fields: fields{cardDeck: CardDeck{cards: []Card(nil)}}, expectError: true},
	}
		for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := tt.fields.cardDeck
			prev := d.DeepCopy()
			assert.EqualValues(t, prev, d)
			if err := d.Shuffle(); (err != nil ) != tt.expectError {
				assert.Fail(t, err.Error())
			} else if !tt.expectError {
				assert.NotEqualValues(t, prev, d)
			}
		})
	}
}