package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	easypoker "github.com/sean9999/EasyPoker"
	"github.com/sean9999/GoCards/french"
)

func main() {

	//	set up
	randy := rand.NewSource(time.Now().UnixMicro())
	d := french.NewShuffledDeck(randy)
	g := easypoker.NewGame(d)
	bob := easypoker.NewPlayer("Bob")
	sally := easypoker.NewPlayer("Sally")

	//	play
	for i := 0; i <= 7; i++ {
		fmt.Printf("\n*****\tRound %d\t*****\n", i)
		r, err := g.NewRound(&bob, &sally)

		if errors.Is(err, easypoker.NotEnoughCardsLeft) {
			fmt.Println(err)
			break
		} else {
			winningHand := r.Play()
			fmt.Println(winningHand.Cards, winningHand.Player.Name)
		}

	}

}
