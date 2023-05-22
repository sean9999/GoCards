package french

import "fmt"

func LookupFace(n Face) (string, error) {
	word, exists := faceMap[Face(n)]
	if exists {
		return word, nil
	}
	return "IllegalValue", fmt.Errorf("illegal Face value %q", n)
}

var faceMap = map[Face]string{
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

// Face is an offset from [SuitRange.LowerBound]
type Face rune

const (
	Ace Face = iota
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

func (f Face) String() string {
	name, _ := LookupFace(f)
	return name
}

func GetFace(c Card) Face {
	//	nice optimisation
	return Face(c%16 - 1)
}
