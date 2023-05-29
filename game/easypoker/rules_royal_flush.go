package easypoker

import (
	"github.com/sean9999/GoCards/deck/french"
	"github.com/sean9999/GoFunctional/fslice"
)

var IsRoyalFlush PokerPattern = PokerPattern{
	Declaration: "is royal flush",
	Description: "cards are in sequence beginning at 10, and of the same suit",
	Grade:       RoyalFlush,
	Func:        fn_IsRoyalFlush,
}

var fn_IsRoyalFlush PatternFunc = func(cards Cards) (bool, PokerHand) {
	ok := false
	ph := ZeroPokerHand

	isRoyal := func(c Card, _ int, _ []Card) bool {
		cRank, _ := c.Rank()
		cardIsRoyal := false
		switch cRank {
		case Rank(french.Ten),
			Rank(french.Jack),
			Rank(french.Queen),
			Rank(french.King),
			Rank(french.Ace):
			cardIsRoyal = true
		default:
			cardIsRoyal = false
		}
		return cardIsRoyal
	}
	isSameSuit := func(c Card, _ int, hand []Card) bool {
		firstSuit, _ := hand[0].Suit()
		thisSuit, _ := c.Suit()
		return thisSuit == firstSuit
	}
	fCards := fslice.From(cards)
	ok = fCards.Every(isRoyal) && fCards.Every(isSameSuit)
	if ok {
		ph.Good = cards.AsLowToHigh()
		ph.Remaining = []Card{}
		ph.Grade = RoyalFlush
	}
	return ok, ph
}
