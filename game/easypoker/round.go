package easypoker

import "fmt"

// A Round is over as soon as it starts. A hand goes to every player. One of them wins
type Round struct {
	Hands       Hands
	WinningHand *Hand
}

func (g *Game) NewRound(players ...*Player) (*Round, error) {
	//	sad path
	totalDesiredDraw := len(players) * 5
	cardsInStock := len(g.Stock)
	if cardsInStock < totalDesiredDraw {
		return nil, fmt.Errorf("%w (%d / %d)", ErrTooFewCards, cardsInStock, totalDesiredDraw)
	}
	if g.PlayState == Played {
		return nil, fmt.Errorf("game PlayState is Played. Can't create a new round")
	}

	//	happy path
	g.PlayState = Playing
	hands := []Hand{}
	for _, p := range players {
		theseCards, _ := g.Stock.Draw(5)
		thisHand := Hand{
			Player: p,
			Cards:  theseCards,
		}
		hands = append(hands, thisHand)
	}
	round := Round{
		Hands:       hands,
		WinningHand: nil,
	}
	g.Rounds = append(g.Rounds, round)
	return &round, nil
}

// plays the round and returns the winning hand
func (r *Round) Play() Hand {
	//	naive play. the first hand wins
	wHand := r.Hands[0]
	r.WinningHand = &wHand
	return wHand
}

func (r Round) Winner() *Player {
	if r.WinningHand == nil {
		return nil
	}
	return r.WinningHand.Player
}
