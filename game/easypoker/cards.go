package easypoker

import (
	"fmt"

	"github.com/sean9999/GoCards/deck/french"
)

type Card french.Card

func (c Card) Beats(d Card) bool {
	return c < d
}

func (c Card) String() string {
	f := french.Card(c)
	return f.String()
}

type Cards []Card

func (cs Cards) Less(i, j int) bool {
	return cs[j].Beats(cs[i])
}
func (cs Cards) Beats(ds Cards) bool {
	//	naive algo: sum the value of each card in the hand
	c_sum := int32(0)
	d_sum := int32(0)
	for _, c := range cs {
		c_sum += int32(c)
	}
	for _, d := range ds {
		d_sum += int32(d)
	}
	return c_sum > d_sum
}

type Stock Cards

func (p *Stock) Draw(n int) (Cards, error) {
	s := *p
	chop := len(s) - n
	if chop < 1 {
		return nil, fmt.Errorf("%d in stock, but wanted %d. %w", len(s), n, ErrTooFewCards)
	} else {
		//	chop tail off
		tail := s[chop:]
		head := s[:chop]
		*p = head
		return Cards(tail), nil
	}
}

// ConstructCard constructs a card from well-known suit and rank values from the [french] deck
func ConstructCard(s french.Suit, r french.Rank) Card {
	val := rune(s.Range().LowerBound) + rune(r)
	return Card(val)
}

// CardFromValue creates a card from a raw rune value
func CardFromValue(v rune) Card {
	return Card(v)
}

// PokerCard takes a french card (a card from the standard french suited deck) and returns the equivalent easypoker card
// the difference is that that easypoker has rules on what card outranks another (ie: aces high)
// or what group of cards (ie: hand) beats another. (ie: full-house beats a pair)
// or what cards are legal (jokers not permitted)
func PokerCard(f french.Card) Card {
	return Card(f)
}

func (c Card) Validate() (bool, error) {
	// if it's not a valid card from the french deck,
	// it's certainly not valid here
	f := french.Card(c)
	okAsFrench, err := f.Validate()
	if !okAsFrench {
		return false, err
	}
	//	jokers to boot
	rank := f.Rank()
	if rank == french.Joker {
		return false, fmt.Errorf("jokers not allowed in Easy Poker")
	}
	return true, nil
}
