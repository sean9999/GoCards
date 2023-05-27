package french

import (
	"math/rand"
)

type Cards []Card

/*
func (c Cards) Beats(d Cards) bool {
	return c[0] > d[0]
}
*/

func StreamCards(randy rand.Source, doneChan <-chan bool) <-chan Card {
	//	cards drawn from randomly shuffled decks
	ch := make(chan Card)
	pool := make([]Card, 0, 54)
	doneVal := false

	// pop off the top card from a shuffled deck
	// get a new deck if we run out
	// this is more realistic than doing pure random card choices
	var poolPop = func() Card {
		if len(pool) < 1 {
			pool = append(pool, NewShuffledDeck(randy).DealOut()...)
		}
		//	pop off the last element
		popeye := pool[len(pool)-1]
		pool = pool[:len(pool)-1]
		return popeye
	}

	go func() {
		for !doneVal {
			select {
			case doneVal = <-doneChan:
				close(ch)
			case ch <- poolPop():
				//	stream out cards as fast as our receiver can take them
			}
		}
	}()

	return ch
}
