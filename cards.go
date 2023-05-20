package cards

import (
	"fmt"
)

type Suit uint8

const (
	IllegalSuit Suit = iota // zero-value should be illegal to protect against accidental values
	Diamonds
	Clubs
	Hearts
	Spades
	NoSuit // Jokers have no suit
)

type Face uint8

const (
	IllegalFace Face = iota // zero-value is illegal
	Ace
	Two // face-values correspond to integer values, allowing easy calculation
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
	Joker
)

type Card struct {
	Suit
	Face
}

type CardException struct {
	Card
}

func NewCardException(c Card) CardException {
	e := CardException{c}
	return e
}

func (c CardException) Error() string {
	if uint8(c.Face)*uint8(c.Suit) == 0 {
		return fmt.Sprintf("Out of bounds Face and Suit %q, %q", c.Face, c.Suit)
	}
	if c.Face < 1 || c.Face > Joker {
		return fmt.Sprintf("Out of bounds Face %q", c.Face)
	}
	if c.Suit < 1 || c.Suit > 5 {
		return fmt.Sprintf("Out of bounds Suit %q", c.Suit)
	}
	if c.Face == Joker {
		if c.Suit != NoSuit {
			return fmt.Sprintf("a Joker cannot be of Suit %q", c.Suit)
		}
	}
	return "Unknown Card Error"
}

func (c Card) Validate() (bool, error) {

	return true, nil
}

func (c Card) String() string {

}

type Deck [56]Card

// Pile is a generic type that can represent a talon, a hand, a layout, a discard pile, etc
type Pile []Card

func NewDeck() Deck {
	d := Deck{}

	i := 0

	for thisSuit := Suit(1); thisSuit <= Spades; thisSuit++ {
		for thisFace := Face(1); thisFace <= Ace; thisFace++ {
			thisCard := Card{
				Suit: thisSuit,
				Face: thisFace,
			}
			d[i] = thisCard
			i++
		}
	}
	return d
}
