package easypoker

import (
	"fmt"
	"sort"

	"github.com/sean9999/GoCards/deck/french"
)

// Cards is a distinct datatype, because we need to compare groups of cards
// in a different way than we compare individual ones
// for example, three threes beats a high ace
type Cards []Card

// sort cards by rank and suit from low to high, as if comparing two single cards
func (cards Cards) SortLowToHigh() {
	sort.Sort(cards)
}

// sort cards by rank and suit from high to low, as if comparing two single cards
func (cards Cards) SortHighToLow() {
	sort.Sort(sort.Reverse(cards))
}

func (cards Cards) AsLowToHigh() Cards {
	newHand := Cards(cards)
	newHand.SortLowToHigh()
	return newHand
}

func (cards Cards) AsHighToLow() Cards {
	newHand := Cards(cards)
	newHand.SortHighToLow()
	return newHand
}

func (cards Cards) HighCard() Card {
	r := ZeroCard
	for _, c := range cards {
		if c > r {
			r = c
		}
	}
	return r
}

func CardsFromFrench(fs french.Cards) (Cards, error) {
	r := make([]Card, 0, len(fs))
	for _, fc := range fs {
		c, err := CardFromFrench(fc)
		if err != nil {
			return r, fmt.Errorf("a joker was encountered")
		}
		r = append(r, c)
	}
	return r, nil
}
