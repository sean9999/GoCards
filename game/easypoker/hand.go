package easypoker

import french "github.com/sean9999/GoCards/french"

// a Hand is five Cards held by a Player
type Hand struct {
	Cards  []french.Card
	Player *Player
}
