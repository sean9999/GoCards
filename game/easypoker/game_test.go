package easypoker_test

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/sean9999/GoCards/deck/french"
	"github.com/sean9999/GoCards/game/easypoker"
)

func assertScalarEquals[C comparable](t *testing.T, got, want C) {
	t.Helper()
	if want != got {
		t.Errorf("want %v but got %v", want, got)
	}
}

func CardsStrings(cs []easypoker.Card) []string {
	r := make([]string, 0, len(cs))
	for _, c := range cs {
		r = append(r, french.Card(c).String())
	}
	return r
}

func cardsAreEqual(t *testing.T, got, want []easypoker.Card) (bool, error) {
	t.Helper()
	if len(got) != len(want) {
		return false, fmt.Errorf("got %v but wanted %v", got, want)
	}
	for i := 0; i < len(got); i++ {
		if got[i] != want[i] {
			return false, fmt.Errorf("got %s but wanted %s", easypoker.Cards(got).Strand(), easypoker.Cards(want).Strand())
		}
	}
	return true, nil
}

func TestNewGame(t *testing.T) {

	//	let's try a helper in scope so we don't need to pass *testing.T

	t.Run("new game, seed value 0", func(t *testing.T) {

		//	deterministic game
		g := easypoker.NewGame(rand.NewSource(0))

		t.Run("number of rounds is zero", func(t *testing.T) {
			assertScalarEquals[int](t, len(g.Rounds), 0)
		})

		t.Run("playState is Unplayed", func(t *testing.T) {
			assertScalarEquals[easypoker.PlayState](t, g.PlayState, easypoker.Unplayed)
		})

		t.Run("number of cards in stock is 52", func(t *testing.T) {
			want := 52
			got := len(g.Stock)
			assertScalarEquals[int](t, got, want)
		})

		player_1 := easypoker.NewPlayer("Alice")
		player_2 := easypoker.NewPlayer("Bob")
		alice := &player_1
		bob := &player_2
		round, _ := g.NewRound(alice, bob)

		t.Run("after new round creation", func(t *testing.T) {
			t.Run("number of rounds is 1", func(t *testing.T) {
				assertScalarEquals[int](t, len(g.Rounds), 1)
			})
			t.Run("that round has no winner", func(t *testing.T) {
				assertScalarEquals[*easypoker.Player](t, round.Winner(), nil)
			})
			t.Run("playState is now Playing", func(t *testing.T) {
				assertScalarEquals[easypoker.PlayState](t, g.PlayState, easypoker.Playing)
			})
		})

		t.Run("after first round is played", func(t *testing.T) {

			winningHand := round.Play()

			t.Run("the winner is Alice", func(t *testing.T) {
				assertScalarEquals[*easypoker.Player](t, winningHand.Player, alice)
			})

			t.Run("returned value is same as recorded value", func(t *testing.T) {

				returnedValue := winningHand
				recordedValue := round.WinningHand

				t.Run("player is same", func(t *testing.T) {
					got := returnedValue.Player
					want := recordedValue.Player
					assertScalarEquals[*easypoker.Player](t, got, want)
				})

				t.Run("hand is same", func(t *testing.T) {
					got := returnedValue.Cards
					want := recordedValue.Cards
					ok, err := cardsAreEqual(t, got, want)
					if !ok {
						t.Error(err)
					}
				})

			})

		})

		t.Run("play a second round", func(t *testing.T) {

			round, err := g.NewRound(alice, bob)
			winningHand := round.Play()
			if err != nil {
				t.Error(err)
			}

			t.Run("the winner is again Alice", func(t *testing.T) {
				assertScalarEquals[*easypoker.Player](t, winningHand.Player, alice)
			})

			t.Run("so the loser must be Bob", func(t *testing.T) {

				var got *easypoker.Player
				want := bob

				for _, h := range round.Hands {
					if h.Player != winningHand.Player {
						got = h.Player
					}
				}

				assertScalarEquals[*easypoker.Player](t, got, want)
			})

			t.Run("32 cards left in stock", func(t *testing.T) {
				assertScalarEquals[int](t, len(g.Stock), 32)
			})

			t.Run("losing hand is 🃕🃂🃃🂡🃄", func(t *testing.T) {
				var losingHand easypoker.Hand
				if round.Hands[0].Player == round.WinningHand.Player {
					losingHand = round.Hands[1]
				} else {
					losingHand = round.Hands[0]
				}
				want, _ := easypoker.Strand("🃕🃂🃃🂡🃄")
				got := losingHand.Cards
				equal, err := cardsAreEqual(t, got, want)
				if !equal {
					t.Error(err)
				}
			})

			t.Run("winning hand is 🃙🂽🂭🃊🃎", func(t *testing.T) {
				want, err := easypoker.Strand("🃙🂽🂭🃊🃎")
				if err != nil {
					t.Error(err)
				}
				got := round.WinningHand.Cards
				equal, err := cardsAreEqual(t, got, want)
				if !equal {
					t.Error(err)
				}
			})

			t.Run("playState should still be playing", func(t *testing.T) {
				want := easypoker.Playing
				got := g.PlayState
				assertScalarEquals[easypoker.PlayState](t, got, want)
			})

		})

		t.Run("third round", func(t *testing.T) {

			round, err := g.NewRound(alice, bob)
			winningHand := round.Play()
			if err != nil {
				t.Error(err)
			}

			t.Run("the winner is again Alice", func(t *testing.T) {
				alice := &player_1
				assertScalarEquals[*easypoker.Player](t, winningHand.Player, alice)
			})
			t.Run("22 cards left in stock", func(t *testing.T) {
				assertScalarEquals[int](t, len(g.Stock), 22)
			})

		})

		t.Run("fourth round", func(t *testing.T) {

			round, err := g.NewRound(alice, bob)
			winningHand := round.Play()
			if err != nil {
				t.Error(err)
			}

			t.Run("the winner is again Alice", func(t *testing.T) {
				alice := &player_1
				assertScalarEquals[*easypoker.Player](t, winningHand.Player, alice)
			})
			t.Run("12 cards left in stock", func(t *testing.T) {
				assertScalarEquals[int](t, len(g.Stock), 12)
			})

		})

	})

}
