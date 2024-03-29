package easypoker

import (
	"fmt"
	"strings"

	"github.com/sean9999/GoCards/deck/french"
)

// a Hand is five Cards held by a Player
type Hand struct {
	Cards  Cards
	Player *Player
}

type Hands []Hand

func ConstructHand(suits []french.Suit, ranks []french.Rank) (Cards, error) {
	//	sad path
	if len(suits) == 0 || len(ranks) == 0 {
		return nil, fmt.Errorf("a zero-length slice cannot be used to construct a hand")
	}
	if len(suits) != len(ranks) {
		return nil, fmt.Errorf("equal length slices are necessary to construct a hand")
	}
	// happy path
	cards := []Card{}
	for i := 0; i < len(suits); i++ {
		thisSuit := suits[i]
		thisRank := ranks[i]
		cards = append(cards, ConstructCard(thisSuit, thisRank))
	}
	return Cards(cards), nil
}

func HandFromString(str string) (Cards, error) {
	chars := strings.Split(str, "")
	return ConstructHandFromChars(chars)
}

func ConstructHandFromChars(chars []string) (Cards, error) {
	cards := make([]Card, 0, len(chars))
	for _, char := range chars {
		if char != " " {
			thisFrenchCard, err := french.CardFromChar(char)
			if err != nil {
				return nil, err
			} else {
				thisPokerCard, _ := CardFromFrench(thisFrenchCard)
				isValid, err := thisPokerCard.Validate()
				if !isValid {
					return nil, err
				} else {
					cards = append(cards, thisPokerCard)
				}
			}
		}
	}
	return Cards(cards), nil
}

func (cs Cards) Strand() string {
	r := ""
	for _, c := range cs {
		r += c.String() + " "
	}
	return r
}

func Strand(longString string) (Cards, error) {
	chars := strings.Split(longString, "")
	return ConstructHandFromChars(chars)
}
