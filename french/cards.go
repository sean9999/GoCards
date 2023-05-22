package french

import "math/rand"

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
			d := NewShuffledDeck(randy)
			pool = append(pool, d[:]...)
		}
		poppedCard := pool[len(pool)-1]
		pool = pool[:len(pool)-1]
		return poppedCard
	}

	go func() {
		for !doneVal {
			select {
			case doneVal = <-doneChan:
				close(ch)
			case ch <- poolPop():

			}
		}
	}()

	return ch
}
