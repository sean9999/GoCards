package french

import (
	"fmt"
	"unicode/utf8"
)

// Card is a rune whose value is a UTF-8 character like 🃈 or 🂮
type Card rune

var ZeroCard = Card(0)

func (c Card) Suit() (Suit, error) {
	return GetSuit(c)
}

func (c Card) Rank() Rank {
	return GetRank(c)
}

func (c Card) String() string {
	buf := utf8.AppendRune(nil, rune(c))
	return string(buf)
}

func (c Card) Word() string {
	suit, _ := c.Suit()
	return fmt.Sprintf("%s of %s", c.Rank(), suit.Word())
}

func (c Card) FallsWithin() SuitRange {
	suit, _ := c.Suit()
	r := LegalSuitRanges[suit]
	return r
}

type CardException struct {
	Card
	Message string
}

func (c CardException) Error() string {
	return c.Message
}

func (c Card) Validate() (bool, error) {
	suit, err := GetSuit(c)
	if err != nil {
		return false, err
	}
	offset := c - suit.Range().Floor()
	// King is the highest value because that's where it sits in the UTF-8 table.
	// Particular games implement their own [Card.Beats].
	// see https://www.unicode.org/charts/PDF/U1F0A0.pdf
	if offset > Card(King) {
		return false, CardException{c, fmt.Sprintf("invalid face value %q. Highest legal value is %q", offset, King)}
	}
	return true, nil
}

func MakeCard(rank Rank, suit Suit) Card {
	val := rune(suit.Range().Floor()) + rune(rank)
	return Card(val)
}

func CardFromChar(char string) (Card, error) {
	runeSlice := []rune(char)
	if len(runeSlice) != 1 {
		return ZeroCard, CardException{
			ZeroCard,
			fmt.Sprintf("runeSlice of length 1 expected. got %d", len(runeSlice)),
		}
	}
	val := runeSlice[0]
	return CardFromRune(val)
}

func CardFromRune(v rune) (Card, error) {
	c := Card(v)
	ok, err := c.Validate()
	if !ok {
		return ZeroCard, err
	}
	return c, nil
}
