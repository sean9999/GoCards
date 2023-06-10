package french

import (
	"math/rand"
	"strings"
)

type Cards []Card

func (cs Cards) Strand() string {
	r := ""
	for _, c := range cs {
		r += c.String() + " "
	}
	return r
}

func Strand(longString string) (Cards, error) {
	chars := strings.Split(longString, "")
	return ConstructHandFromChars(chars)
}

func StreamCards(randy rand.Source) (<-chan Card, chan<- bool) {
	//	cards drawn from randomly shuffled decks
	cardsChan := make(chan Card)
	doneChan := make(chan bool)

	//defer close(doneChan)
	//defer close(cardsChan)

	pool := make([]Card, 0, 54)

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
		doneVal := false
		defer close(doneChan)
		defer close(cardsChan)

		for !doneVal {
			select {
			case <-doneChan:
				doneVal = true
			case cardsChan <- poolPop():
				//	stream out cards as fast as our receiver can take them
			}
		}
	}()

	return cardsChan, doneChan
}

func ConstructHandFromChars(chars []string) (Cards, error) {
	cards := make([]Card, 0, len(chars))
	for _, char := range chars {
		if char != " " {
			thisFrenchCard, err := CardFromChar(char)
			if err != nil {
				return nil, err
			} else {
				cards = append(cards, thisFrenchCard)
			}
		}
	}
	return Cards(cards), nil
}
