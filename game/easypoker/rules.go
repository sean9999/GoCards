package easypoker

import (
	"fmt"

	"github.com/sean9999/GoFunctional/fslice"
)

// PatternFunc describes a poker-hand via a function that operates over the cards
type PatternFunc func(cards Cards) (bool, PokerHand)

// constants that can be compared to determine what PokerHand beats another
type PatternGrade uint8

const (
	NoGrade PatternGrade = iota
	HighCard
	Pair
	TwoPair
	ThreeOfAKind
	Straight
	Flush
	FullHouse
	FourOfAKind
	StraightFlush
	RoyalFlush
)

type PokerPattern struct {
	Grade       PatternGrade
	Declaration string
	Description string
	Func        PatternFunc
}

// a hand grouped into the important cards that make up a poker hand, the remainders
// and what kind of poker hand it is (Grade)
type PokerHand struct {
	Good      Cards
	Remaining Cards
	Grade     PatternGrade
}

func HighestPokerHand(cards Cards) PokerHand {
	type patternTuple struct {
		Grade   PatternGrade
		Pattern PokerPattern
	}
	patterns := []patternTuple{
		{RoyalFlush, IsRoyalFlush},
		{StraightFlush, IsStraightFlush},
		{FourOfAKind, HasFourOfAKind},
		{FullHouse, IsFullHouse},
		{Flush, IsFlush},
		{Straight, IsStraight},
		{ThreeOfAKind, HasThreeOfAKind},
		{TwoPair, HasTwoPair},
		{Pair, HasPair},
	}
	for _, tuple := range patterns {
		is, pokerHand := tuple.Pattern.Func(cards)
		if is {
			pokerHand.Grade = tuple.Grade
			return pokerHand
		}
	}
	// if we haven't returned yet
	// our best poker hand is simply a high card
	_, highCardHand := GetHighCard.Func(cards)
	return highCardHand
}

var ZeroPokerHand PokerHand = PokerHand{}

func GradeFactory(n int) (PatternGrade, error) {
	switch n {
	case 2:
		return Pair, nil
	case 3:
		return ThreeOfAKind, nil
	case 4:
		return FourOfAKind, nil
	default:
		err := fmt.Errorf("out of range for PokerHandGrade: %d", n)
		return 0, err
	}
}

func OfAKindFactory(n int) PatternFunc {
	return func(cards Cards) (bool, PokerHand) {
		ok := false
		ph := PokerHand{}

		fCards := fslice.From(cards)
		ok = fCards.Some(func(c Card, _ int, _ []Card) bool {
			cRank, _ := c.Rank()
			cardsMatchingThisRank := fCards.Filter(func(d Card, _ int, _ []Card) bool {
				dRank, _ := d.Rank()
				return cRank == dRank
			})
			if len(cardsMatchingThisRank) == n {
				ph.Good = cardsMatchingThisRank.ToSlice()
				ph.Remaining = fCards.Filter(func(e Card, _ int, _ []Card) bool {
					return !cardsMatchingThisRank.Includes(e)
				}).ToSlice()
				ph.Grade, _ = GradeFactory(n)
				return true
			}
			return false
		})
		return ok, ph
	}
}
