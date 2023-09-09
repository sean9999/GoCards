package easypoker_test

import (
	"math/rand"
	"testing"

	"github.com/sean9999/GoCards/game/easypoker"
)

func TestCardSort(t *testing.T) {

	randy := rand.NewSource(0)
	g := easypoker.NewGame(randy)

	hand1String := "🃆🃓🂹🃞🂮"
	t.Run("first hand is "+hand1String, func(t *testing.T) {
		hand1, _ := g.Stock.Draw(5)
		got := hand1
		want, _ := easypoker.Strand(hand1String)
		equal, err := cardsAreEqual(t, got, want)
		if !equal {
			t.Error(err)
		}

		hand1StringLowToHigh := "🃓🃆🂹🃞🂮"
		t.Run("sorted low to high is "+hand1StringLowToHigh, func(t *testing.T) {
			hand1.SortLowToHigh()
			got := hand1
			want, _ := easypoker.Strand(hand1StringLowToHigh)
			equal, err := cardsAreEqual(t, got, want)
			if !equal {
				t.Error(err)
			}
		})
	})

	hand2String := "🃘🂣🂵🃚🃂"
	t.Run("second hand is "+hand2String, func(t *testing.T) {
		hand2, _ := g.Stock.Draw(5)
		got := hand2
		want, _ := easypoker.Strand(hand2String)
		equal, err := cardsAreEqual(t, got, want)
		if !equal {
			t.Error(err)
		}

		hand2StringHighToLow := "🃚🃘🂵🂣🃂"
		t.Run("sorted high to low is "+hand2StringHighToLow, func(t *testing.T) {
			hand2.SortHighToLow()
			got := hand2
			want, _ := easypoker.Strand(hand2StringHighToLow)
			equal, err := cardsAreEqual(t, got, want)
			if !equal {
				t.Error(err)
			}
		})

	})

}
