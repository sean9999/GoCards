package cards

import "math/rand"

type Game[DECK_TYPE any] struct {
	Deck  DECK_TYPE
	State map[string]any
}

type Deck interface {
	Shuffle(rand.Source)
	Valid() (bool, error)
}

type Card interface {
	Suit() Suit
	Face() Face
	String() string // ex: 🂮
	Word() string   // ex: King of Spades
	Code() string   // ex: K♠
	Valid() (bool, error)
	Value() rune
}

type Suit interface {
	String() string // ex: ♠
	Value() rune
	Word() string // ex: "Spades"
}

// note that although Face is a rune, it does not correspond to a UTF-8 char
// because there is no visual representation of a face without a suit
type Face interface {
	String() // ex: "King"
	Code()   // ex: "K"
}

/*
func StreamCards(seed int64) chan Card {
	//	cards drawn from randomly shuffled decks
	ch := make(chan Card)
	pool := make([]Card, 0, 54)

	randy := NewDeterminator(seed)

	go func() {

		//	if we've exhausted the deck, get a new one
		if len(pool) < 1 {
			d := NewDeck()
			d.Shuffle(randy)
			randy.Tick()
			pool = append(pool, d[:]...)
		}
		//	draw top card from deck
		popped := pool[len(pool)-1]
		pool = pool[:len(pool)-1]
		ch <- popped

	}()

	return ch
}

func StreamDecks(seed int64) chan<- Deck {
	ch := make(chan Deck)
	randy := NewDeterminator(seed)

	go func() {
		d := NewDeck()
		d.Shuffle(randy)
		randy.Tick()
		ch <- d
	}()

	return ch
}


func (d *Deck) Shuffle(randy rand.Source) {
	generator := rand.New(randy)
	generator.Shuffle(len(d), func(i, j int) { d[i], d[j] = d[j], d[i] })
}

func NewShuffledDeck(randy rand.Source) Deck {
	d := NewDeck()
	d.Shuffle(randy)
	return d
}

*/
