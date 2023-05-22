package cards

import "math/rand"

// Pile is a generic type that can represent a talon, a hand, a layout, a discard pile, etc
// A pile is derived from a Deck
type Pile []Card

func (p Pile) Shuffle(seed rand.Source) Pile {
	generator := rand.New(seed)
	generator.Shuffle(len(p), func(i, j int) { p[i], p[j] = p[j], p[i] })
	return p
}
