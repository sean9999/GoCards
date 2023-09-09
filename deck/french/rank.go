package french

import "fmt"

func LookupRank(n Rank) (string, error) {
	word, exists := rankMap[Rank(n)]
	if exists {
		return word, nil
	}
	return "IllegalValue", fmt.Errorf("illegal Face value %q", n)
}

var rankMap = map[Rank]string{
	Ace:    "Ace",
	Two:    "Two",
	Three:  "Three",
	Four:   "Four",
	Five:   "Five",
	Six:    "Six",
	Seven:  "Seven",
	Eight:  "Eight",
	Nine:   "Nine",
	Ten:    "Ten",
	Jack:   "Jack",
	Knight: "Knight",
	Queen:  "Queen",
	King:   "King",
	Joker:  "Joker",
}

// Rank is an offset from [SuitRange.LowerBound]
// we purposely do not make it a rune, because it alone does not represent a UTF-8 char
type Rank uint8

const (
	ZeroRank Rank = iota
	Ace
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Knight
	Queen
	King
	Joker
)

func (f Rank) String() string {
	name, _ := LookupRank(f)
	return name
}

func GetRank(c Card) Rank {
	//	nice optimisation
	return Rank(c % 16)
}
