package easypoker

import (
	"fmt"

	"github.com/sean9999/GoCards/deck/french"
)

type Rank french.Rank

var ZeroRank Rank = Rank(french.ZeroRank)

// this list specifically excludes Knights and Jokers
var AllowedRanks []french.Rank = []french.Rank{
	french.Two,
	french.Three,
	french.Four,
	french.Five,
	french.Six,
	french.Seven,
	french.Eight,
	french.Nine,
	french.Ten,
	french.Jack,
	french.Queen,
	french.King,
	french.Ace,
}

func (c Card) Rank() (Rank, error) {
	thisRank := french.Card(c).Rank()
	for _, legalRank := range AllowedRanks {
		if thisRank == legalRank {
			return Rank(thisRank), nil
		}
	}
	return ZeroRank, fmt.Errorf("illegal Rank %q", thisRank)
}

func (r Rank) Next() (Rank, error) {
	frenchRank := french.Rank(r)
	switch frenchRank {
	case french.King:
		return Rank(french.Ace), nil
	case french.Ace:
		return ZeroRank, fmt.Errorf("there is no rank higher than Ace")
	case french.Jack:
		// skip over Knight
		return Rank(french.Queen), nil
	default:
		return Rank(r + 1), nil
	}
}

func (r Rank) Previous() (Rank, error) {
	frenchRank := french.Rank(r)
	switch frenchRank {
	case french.Ace:
		return Rank(french.King), nil
	case french.Two:
		return ZeroRank, fmt.Errorf("there is no rank lower than Two")
	case french.Queen:
		return Rank(french.Jack), nil
	default:
		return Rank(r - 1), nil
	}
}
