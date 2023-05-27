package easypoker

import (
	"fmt"
)

type Stock Cards

func (p *Stock) Draw(n int) (Cards, error) {
	s := *p
	chop := len(s) - n
	if chop < 1 {
		return nil, fmt.Errorf("%d in stock, but wanted %d. %w", len(s), n, ErrTooFewCards)
	} else {
		//	chop tail off
		tail := s[chop:]
		head := s[:chop]
		*p = head
		return Cards(tail), nil
	}
}
