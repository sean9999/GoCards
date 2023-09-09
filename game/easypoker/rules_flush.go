package easypoker

import (
	"github.com/sean9999/GoFunctional/fslice"
)

var IsFlush = PokerPattern{
	Declaration: "is a flush",
	Description: "all cards are of the same suit",
	Grade:       Flush,
	Func:        fn_IsFlush,
}

var fn_IsFlush PatternFunc = func(cards Cards) (bool, PokerHand) {
	ok := false
	ph := PokerHand{}

	fCards := fslice.From(cards)
	//	compare the suit of every card to the suit of the first card
	firstSuit, _ := cards[0].Suit()
	ok = fCards.Every(func(c Card, _ int, _ []Card) bool {
		thisSuit, _ := c.Suit()
		return thisSuit == firstSuit
	})
	if ok {
		ph = PokerHand{
			Good:      cards.AsLowToHigh(),
			Remaining: []Card{},
			Grade:     Flush,
		}
	}
	return ok, ph
}
