package french

import "fmt"

type Stock []Card

func (p *Stock) Draw(n int) (Cards, error) {
	s := *p
	if n > len(s) {
		return nil, fmt.Errorf("too few cards in stock")
	}
	tail := s[n:]
	head := s[:n]
	*p = tail
	return Cards(head), nil
}
