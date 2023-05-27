package easypoker

// compare two individual cards
func (c Card) Beats(d Card) bool {
	if c.Rank() != d.Rank() {
		return c.Rank().Beats(d.Rank())
	} else {
		c_suit, _ := c.Suit()
		d_suit, _ := d.Suit()
		return c_suit.Beats(d_suit)
	}
}

func (cs Cards) Beats(ds Cards) bool {
	//	naive algo: holder of the highest hand wins
	c_sum := int32(0)
	d_sum := int32(0)
	for _, c := range cs {
		c_sum += int32(c)
	}
	for _, d := range ds {
		d_sum += int32(d)
	}
	return c_sum > d_sum
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
