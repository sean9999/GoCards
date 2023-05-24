package french

type Stock []Card

func (p *Stock) Draw(n int) Cards {
	s := *p
	chop := len(s) - n
	tail := s[chop:]
	//	chop tail off
	head := s[:chop]
	*p = head
	return Cards(tail)
}

/*
func (c Stock) Cards() []Card {
	r := make([]Card, len(c))
	copy(r, c)
	return r
}
*/
