package easypoker

import (
	"github.com/sean9999/GoCards/deck/french"

	"github.com/sean9999/GoFunctional/fslice"
)

var IsStraight = PokerPattern{
	Declaration: "is a straight",
	Grade:       Straight,
	Func:        fn_IsStraight,
}

var fn_IsStraight PatternFunc = func(cards Cards) (bool, PokerHand) {
	ok := false
	ph := PokerHand{}

	//	sort low to high
	fCards := fslice.From(cards.AsLowToHigh())
	// is every card the last card + 1?
	// use [Rank.Previous] to evaluate this condition
	ok = fCards.Every(func(c Card, i int, hand []Card) bool {
		thisRank, _ := c.Rank()
		if i == 0 {
			//	if the lowest card is a jack or higher, a straight is impossible
			return thisRank < Rank(french.Jack)
		}
		lastCard := hand[i-1]
		lastRank, _ := lastCard.Rank()         // the rank of the last card
		previousRank, _ := thisRank.Previous() // exactly one step down from this card's rank
		return lastRank == previousRank
	})
	if ok {
		ph.Good = fCards.ToSlice()
		ph.Remaining = []Card{}
		ph.Grade = RoyalFlush
	}
	return ok, ph
}
