package cards

import "math/rand"

type PlayState uint8

const (
	Unplayed PlayState = iota
	Playing
	Played
)

type Game interface {
	Results() any
	Rounds() []Round
	PlayState() PlayState
	NewRound(...Player) (Round, error) // if a new round is impossible, returns an error
}

type Player interface {
	Hand() Cards
	PlayState() PlayState
}

// Round is a group of Hands associated with players
// Every round has a winner or a tie
type Round interface {
	Play() (Player, error) // makes every player play. returns the winner
	Hands() []Hand
	PlayState() PlayState
}

// a Deck is all the cards that could be played, representing a full complete deck
type Deck interface {
	Shuffle(rand.Source)
	DealOut() Stock
}

// Stock is the pile of unplayed cards. It is derived from a deck at the beginning of a game
type Stock interface {
	Draw(n int) Cards // Draw from self and return n Cards
	Cards() []Card
}

// a Hand is a set of cards in a player's hand
type Hand interface {
	Player() Player
	Cards() []Card
}

// Cards is a group of cards. It could be a pile, a hand, a discard pile, etc
// it's conceptually similar to a stock, but we require a different set of methods
// Cards.Beats evaluates groups of cards and decides which is the better hand.
// It can compare Cards ([]Card) of different lengths
// ie: does an Ace of Spades beat two threes? Does an Ace of Hearts beat a 5 and a 7?
type Cards interface {
	Beats(Cards) bool
	Cards() []Card
}

// Card is a playing card
type Card interface {
	Suit() Suit
	Rank() Rank
	String() string // ie: 🂮
	Beats(Card) bool
}

type Suit interface {
	String() string // ie: ♠
	Beats(Suit) bool
}

// Rank is the face-value of a card, irrespective of it's suit
type Rank interface {
	String() string // ex: "King"
	Beats(Rank) bool
}
