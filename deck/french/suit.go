package french

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

var ErrOutofRange = errors.New("suit out of range")

type Suit rune

const (
	ZeroSuit Suit = 0 // should be illegal
	Clubs    Suit = 0x2663
	Diamonds Suit = 0x2666
	Hearts   Suit = 0x2665
	Spades   Suit = 0x2660
	Black    Suit = 0x1F0CF // Black Joker
	Red      Suit = 0x1F0BF // Red Joker
	White    Suit = 0x1F0DF // White Joker
)

// SuitRange is an _inclusive_ range of values to determine a card's suit
type SuitRange struct {
	LowerBound Card // inclusive
	UpperBound Card // inclusive
}

// Floor represents the _exclusive_ lower-bound
// useful in preventing off-by-one errors
// when combining Suit and Rank values to calculate Card's underlying rune value
func (sr SuitRange) Floor() Card {
	return sr.LowerBound - 1
}

// the UTF8 character representing the suit
func (s Suit) String() string {
	buf := utf8.AppendRune(nil, rune(s))
	return string(buf)
}

// the range of runes representing all cards in this suit
func (s Suit) Range() SuitRange {
	return LegalSuitRanges[s]
}

func (s Suit) Validate() (bool, error) {
	_, ok := LegalSuitRanges[s]
	if !ok {
		return false, fmt.Errorf("%w: %q", ErrOutofRange, s)
	}
	return true, nil
}

// the suit as an english word
func (s Suit) Word() string {
	switch s {
	case Diamonds:
		return "Diamonds"
	case Clubs:
		return "Clubs"
	case Hearts:
		return "Hearts"
	case Spades:
		return "Spades"
	case Black:
		return "Black"
	case Red:
		return "Red"
	case White:
		return "White"
	}
	return fmt.Sprintf("Illegal Suit %d", s)
}

var ZeroRange = SuitRange{0, 0}
var SpadesRange = SuitRange{0x1F0A1, 0x1F0AE}
var HeartsRange = SuitRange{0x1F0B1, 0x1F0BE}
var DiamondsRange = SuitRange{0x1F0C1, 0x1F0CE}
var ClubsRange = SuitRange{0x1F0D1, 0x1F0DE}
var BlackJokerRange = SuitRange{0x1F0CF, 0x1F0CF}
var RedJokerRange = SuitRange{0x1F0BF, 0x1F0BF}
var WhiteJokerRange = SuitRange{0x1F0DF, 0x1F0DF}

var LegalSuitRanges = map[Suit]SuitRange{
	Spades:   SpadesRange,
	Hearts:   HeartsRange,
	Diamonds: DiamondsRange,
	Clubs:    ClubsRange,
	Black:    BlackJokerRange,
	Red:      RedJokerRange,
	White:    WhiteJokerRange,
}

func cardFallsWithinRange(c Card, r SuitRange) bool {
	if c >= r.LowerBound && c <= r.UpperBound {
		return true
	}
	return false
}

func GetSuit(c Card) (Suit, error) {
	if c == ZeroCard {
		return ZeroSuit, CardException{c, "Zero card has no legal suit"}
	}
	for thisSuit, thisRange := range LegalSuitRanges {
		if cardFallsWithinRange(c, thisRange) {
			return thisSuit, nil
		}
	}
	return ZeroSuit, CardException{c, "no legal suit for this card"}
}

func (s1 Suit) Beats(s2 Suit) bool {

	s := map[Suit]uint8{
		ZeroSuit: 0,
		Clubs:    1,
		Diamonds: 2,
		Hearts:   3,
		Spades:   4,
	}

	return s[s1] > s[s2]
}
