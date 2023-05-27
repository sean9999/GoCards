package easypoker

import "github.com/sean9999/GoCards/deck/french"

type Rank french.Rank

func (c Card) Rank() Rank {
	return Rank(french.Card(c).Rank())
}
