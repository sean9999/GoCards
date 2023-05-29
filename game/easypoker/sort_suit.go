package easypoker

import "github.com/sean9999/GoCards/deck/french"

func (s1 Suit) Beats(s2 Suit) bool {
	suitRanking := map[french.Suit]uint8{
		french.Clubs:    1,
		french.Diamonds: 2,
		french.Hearts:   3,
		french.Spades:   4,
	}
	thisRank := suitRanking[french.Suit(s1)]
	otherRank := suitRanking[french.Suit(s2)]
	return thisRank > otherRank
}
