package french

import "math/rand"

// Deck is a Standard 52-card deck of French-suited playing cards
//
//	https://en.wikipedia.org/wiki/Standard_52-card_deck
type Deck [54]Card

func NewDeck() Deck {
	var pile []Card
	for thisSuit, thisSuitRange := range LegalSuitRanges {
		for cardValue := thisSuitRange.LowerBound; cardValue <= thisSuitRange.UpperBound; cardValue++ {
			switch thisSuit {
			case Clubs, Hearts, Spades, Diamonds:
				//	Regular cards
				if cardValue-thisSuitRange.LowerBound != Card(Knight) {
					//	There is no knight in a french deck
					pile = append(pile, cardValue)
				}
			case Red, Black:
				//	Jokers
				pile = append(pile, cardValue)
			}
		}
	}
	var d Deck
	copy(d[:], pile)
	return d
}

func (d *Deck) Shuffle(randy rand.Source) {
	generator := rand.New(randy)
	generator.Shuffle(len(d), func(i, j int) { d[i], d[j] = d[j], d[i] })
}

func NewShuffledDeck(randy rand.Source) Deck {
	d := NewDeck()
	d.Shuffle(randy)
	return d
}
