package cards

import "math/rand"

type IGame[DECK_TYPE any] struct {
	Deck  DECK_TYPE
	State map[string]any
}

// a Deck is all the cards that could be played
type IDeck interface {
	Shuffle(rand.Source)
	Valid() (bool, error)
}

// a Stock is all the cards that will be played
type IStock interface {
	Shuffle(rand.Source)
	Valid() (bool, error)
}

// a Hand can be sorted as poker players do
type IHand interface {
	Valid() (bool, error)
	Less(i, j int) bool
}

// Hands ( []Hand ) can be sorted to find the winner and loser
type IHands interface {
	Less(i, j int) bool
}

// Card is a playing card
type ICard interface {
	Suit() ISuit
	Face() IFace
	String() string // ex: 🂮
	Word() string   // ex: King of Spades
	Code() string   // ex: K♠
	Valid() (bool, error)
	Value() rune
}

type ISuit interface {
	String() string // ex: ♠
	Value() rune
	Word() string // ex: "Spades"
}

// note that although Face is a rune, it does not correspond to a UTF-8 char
// because there is no visual representation of a face without a suit
// rather, it represents on offset from the lowest card in the suit to the highest
// as such, it could be a byte rather than a rune, but we leave it to avoid casting
type IFace interface {
	String() // ex: "King"
	Code()   // ex: "K"
}
