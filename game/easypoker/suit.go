package easypoker

import "github.com/sean9999/GoCards/deck/french"

type Suit french.Suit

func (s Suit) String() string {
	return french.Suit(s).String()
}

func (s Suit) Validate() (bool, error) {
	return french.Suit(s).Validate()
}

func GetSuit(c Card) (Suit, error) {
	s, err := french.Card(c).Suit()
	if err != nil {
		return Suit(0), err
	}
	return Suit(s), nil
}
