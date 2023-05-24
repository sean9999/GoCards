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
	if offset > Card(King) {
		return false, CardException{c, fmt.Sprintf("invalid face value %q. Highest legal value is %q", offset, King)}
	}
	//	happy path
	return true, nil
}

/*
func CreateCard(s Suit, f Face) {
	c := Card{s, f}
}
*/
