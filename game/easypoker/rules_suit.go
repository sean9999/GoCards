package easypoker

import "github.com/sean9999/GoCards/deck/french"

func (s1 Suit) Beats(s2 Suit) bool {
	suitRanking := map[french.Suit]int{
		french.Clubs:    1,
		french.Diamonds: 2,
		french.Hearts:   3,
		french.Spades:   4,
	}
	s1Rank := suitRanking[french.Suit(s1)]
	s2Rank := suitRanking[french.Suit(s2)]
	return s1Rank > s2Rank
}
