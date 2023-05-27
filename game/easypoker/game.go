package easypoker

import (
	"errors"
	"math/rand"

	french "github.com/sean9999/GoCards/deck/french" // french-suited standard deck
)

var ErrTooFewCards = errors.New("not enough cards in stock")

type PlayState uint8

const (
	Unplayed PlayState = iota
	Playing
	Played
)

// Game is a simple game of 5-card poker.
// No draws. No wild cards. No folding or betting.
// Players simply play the cards they're dealt, and one of them wins
// A game ends when the deck is exhausted
type Game struct {
	Stock     Stock
	Rounds    []Round
	PlayState PlayState
}

// Winner of a game is they who won most rounds
func (g Game) Winner() *Player {
	var winner *Player
	scores := make(map[*Player]int, len(g.Rounds))
	for _, round := range g.Rounds {
		scores[round.Winner()]++
	}
	highScore := 0
	for player, score := range scores {
		if score > highScore {
			winner = player
			highScore = score
		}
	}
	return winner
}

// get a standard 52-card french-suited deck
// take out the jokers. Easy Poker doesn't use them
func NewGame(randy rand.Source) Game {
	deck := french.NewShuffledDeck(randy)
	stock := make([]Card, 0, 52)
	for _, card := range deck {
		if card.Rank() != french.Joker {
			stock = append(stock, PokerCard(card))
		}
	}
	g := Game{
		Stock:  stock,
		Rounds: []Round{},
	}
	return g
}

func NewDeterministicGame() Game {
	deck := french.NewDeck()
	stock := make([]Card, 0, 52)
	for _, card := range deck {
		if card.Rank() != french.Joker {
			stock = append(stock, PokerCard(card))
		}
	}
	g := Game{
		Stock:  stock,
		Rounds: []Round{},
	}
	return g
}
