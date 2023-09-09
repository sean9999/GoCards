package easypoker

// compare two individual cards
func (c Card) Beats(d Card) bool {
	cRank, _ := c.Rank()
	dRank, _ := d.Rank()
	if cRank != dRank {
		return cRank.Beats(dRank)
	} else {
		c_suit, _ := c.Suit()
		d_suit, _ := d.Suit()
		return c_suit.Beats(d_suit)
	}
}
