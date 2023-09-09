package easypoker

import (
	"fmt"

	"github.com/sean9999/GoCards/deck/french"
)

type Card french.Card

var ZeroCard Card = Card(french.ZeroCard)

func (c Card) Suit() (Suit, error) {
	return GetSuit(c)
}

func (c Card) String() string {
	f := french.Card(c)
	return f.String()
}

// ConstructCard constructs a card from well-known suit and rank values from the [french] deck
func ConstructCard(s french.Suit, r french.Rank) Card {
	val := rune(s.Range().Floor()) + rune(r)
	return Card(val)
}

// CardFromRune creates a card from a raw rune value
func CardFromRune(v rune) Card {
	return Card(v)
}

// CardFromFrench takes a french card (a card from the standard french suited deck) and returns the equivalent easypoker card.
// The difference is that that easypoker has rules on what card outranks another (ie: aces high)
// or what group of cards (ie: hand) beats another. (ie: full-house beats a pair)
// or what cards are legal (jokers are not).
func CardFromFrench(f french.Card) (Card, error) {
	if f.Rank() == french.Joker {
		return ZeroCard, fmt.Errorf("EasyPoker has no Jokers")
	}
	return Card(f), nil
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
