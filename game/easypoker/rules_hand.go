package easypoker

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
