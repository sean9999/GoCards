package easypoker

import "github.com/sean9999/GoCards/deck/french"

func (r Rank) Beats(q Rank) bool {
	//	make aces high rather than low
	if french.Rank(r) == french.Ace {
		r = 16
	}
	if french.Rank(q) == french.Ace {
		q = 16
	}
	return r > q
}
