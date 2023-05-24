package easypoker

import (
	"errors"
	"fmt"

	french "github.com/sean9999/GoCards/deck/french" // french-suited standard deck
)

var ErrTooFewCards = errors.New("not enough cards in stock")

// Game is a simple game of 5-card poker.
// No draws. No wild cards. No folding or betting.
// Players simply play the cards they're dealt, and one of them wins
// A game ends when the deck is exhausted
type Game struct {
	Stock  []french.Card
	Rounds []Round
}

// Draw from the Stock pile. Usually 5, to give to a player
func (g *Game) Draw(n int) ([]french.Card, error) {
	// @todo: error checking if stock is too low
	if len(g.Stock) < n {
		return nil, fmt.Errorf("%d in stock, but wanted %d. %w", len(g.Stock), n, ErrTooFewCards)
	}

	s := g.Stock
	chop := len(s) - n
	tail := s[chop:]
	g.Stock = s[:chop]
	return tail, nil
}

func NewGame(deck french.Deck) Game {
	stock := make([]french.Card, 0, 52)
	for _, card := range deck {
		if card.Face() != french.Joker {
			stock = append(stock, card)
		}
	}
	g := Game{
		Stock:  stock,
		Rounds: []Round{},
	}
	return g
}
