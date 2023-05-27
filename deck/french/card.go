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

func (c Card) Beats(d Card) bool {
	if c.Rank() != d.Rank() {
		return c.Rank().Beats(d.Rank())
	} else {
		c_suit, _ := c.Suit()
		d_suit, _ := d.Suit()
		return c_suit.Beats(d_suit)
	}
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
	ranges := []SuitRange{SpadesRange, HeartsRange, DiamondsRange, ClubsRange}
	for _, thisRange := range ranges {
		return thisRange
	}
	return ZeroRange
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
	offset := c - suit.Range().LowerBound
	// King is the highest value because that's where it sits in the UTF-8 table.
	// This is not a judgement on what card is worth more in a particular game
	// which games implement using the `Card.Beats(Card) bool` method
	// This method only determines the card is in the right range of values according to:
	// https://www.unicode.org/charts/PDF/U1F0A0.pdf
	if offset > Card(King) {
		return false, CardException{c, fmt.Sprintf("invalid face value %q. Highest legal value is %q", offset, King)}
	}
	//	happy path
	return true, nil
}

func NewCard(s Suit, r Rank) Card {
	val := rune(s.Range().LowerBound) + rune(r)
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
	return Card(val), nil
}
