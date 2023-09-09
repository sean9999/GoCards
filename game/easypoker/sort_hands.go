package easypoker

// compares two [Cards], which underneath is just []Card
// colloquially, this compares two "hands"
// But [Hands] means something slightly different in EasyPoker.
// Hands is a Cards with a *Player. [Player] is not actually needed to see if a hand beats another
// Therefore, the meat of the logic is here.
// [Hand.Beats] also exists, but merely wraps this
func (thisHand Cards) Beats(thatHand Cards) bool {
	// four of kind > three of a kind > pair, etc
	thisPokerHand := HighestPokerHand(thisHand)
	thatPokerHand := HighestPokerHand(thatHand)
	if thisPokerHand.Grade > thatPokerHand.Grade {
		return true
	}
	if thisPokerHand.Grade < thatPokerHand.Grade {
		return false
	}
	//	if poker hands were the same, compare high-cards
	thisHighCard := thisPokerHand.Good.HighCard()
	thatHighCard := thatPokerHand.Good.HighCard()
	return thisHighCard > thatHighCard
}

// sort a group of cards from lowest to highest
func (cs Cards) Less(i, j int) bool {
	return cs[j].Beats(cs[i])
}
func (cs Cards) Len() int {
	return len(cs)
}
func (cs Cards) Swap(i, j int) {
	cs[i], cs[j] = cs[j], cs[i]
}

// Since a hand can beat another hand, a slice of hands can be sorted
func (hands Hands) Len() int {
	return len(hands)
}
func (hands Hands) Swap(i, j int) {
	hands[i], hands[j] = hands[j], hands[i]
}
func (hands Hands) Less(i, j int) bool {
	return hands[i].Beats(hands[j])
}

// a Hand beats a Hand if it's underlying Cards beats the other hand's underlying Cards
func (h Hand) Beats(j Hand) bool {
	return h.Cards.Beats(j.Cards)
}
