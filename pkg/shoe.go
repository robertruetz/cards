package pkg

type Shoe struct {
	Deck CardDeck
}

func NewShoe(decks int) *Shoe {
	s := Shoe{
		Deck: CardDeck{},
	}
	for d := 0; d < decks; d++ {
		nd := NewCardDeck()
		s.Deck.Combine(&nd)
	}
	return &s
}

func (s *Shoe) Shuffle() error {
	return s.Deck.Shuffle()
}
