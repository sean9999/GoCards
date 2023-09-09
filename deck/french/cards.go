package french

import (
	"math/rand"
	"strings"
)

type Cards []Card

// Cards.Strand shows a group of cards as a sequence of UTF-8 chars
func (cs Cards) Strand() string {
	r := ""
	for _, c := range cs {
		r += c.String() + " "
	}
	return r
}

// StreamCards streams cards on it's returned channel by producing and exausting shuffled decks card by card.
// You must pass in your done channel and send to it when you want to stop streaming
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

// Strand constructs a Cards from a string
func Strand(longString string) (Cards, error) {
	chars := strings.Split(longString, "")
	return ConstructCardsFromChars(chars)
}

// ConstructCardsFromChars constructs a Cards from a slice of strings
func ConstructCardsFromChars(chars []string) (Cards, error) {
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
