package easypoker_test

import (
	"fmt"
	"testing"

	"github.com/sean9999/GoCards/deck/french"
	"github.com/sean9999/GoCards/game/easypoker"
)

func patternMatchingTest(t *testing.T, handString string, pattern easypoker.PokerPattern, want bool) {
	t.Helper()
	t.Run(pattern.Declaration, func(t *testing.T) {
		t.Helper()
		hand, _ := easypoker.Strand(handString)
		got, _ := pattern.Func(hand)
		if want != got {
			t.Errorf("\"%s %s\" should be a %t statement", handString, pattern.Declaration, want)
		}
	})
}

func TestPokerHands(t *testing.T) {

	var thisHand string

	thisHand = "🃙🂽🂭🃊🃎"
	t.Run(thisHand, func(t *testing.T) {
		patternMatchingTest(t, thisHand, easypoker.HasPair, true)
		patternMatchingTest(t, thisHand, easypoker.HasTwoPair, false)
		patternMatchingTest(t, thisHand, easypoker.HasThreeOfAKind, false)
	})

	thisHand = "🃕🃂🃃🂡🃄"
	t.Run(thisHand, func(t *testing.T) {
		patternMatchingTest(t, thisHand, easypoker.HasPair, false)
		patternMatchingTest(t, thisHand, easypoker.HasThreeOfAKind, false)
		patternMatchingTest(t, thisHand, easypoker.IsStraight, false)
		patternMatchingTest(t, thisHand, easypoker.IsFlush, false)
	})

	thisHand = "🂶🂷🂸🂹🂺"
	t.Run(thisHand, func(t *testing.T) {
		patternMatchingTest(t, thisHand, easypoker.IsStraight, true)
		patternMatchingTest(t, thisHand, easypoker.IsFlush, true)
		patternMatchingTest(t, thisHand, easypoker.IsStraightFlush, true)
		patternMatchingTest(t, thisHand, easypoker.IsRoyalFlush, false)
	})

	thisHand = "🂽🂶🂷🂸🂹"
	t.Run(thisHand, func(t *testing.T) {
		patternMatchingTest(t, thisHand, easypoker.IsStraight, false)
		patternMatchingTest(t, thisHand, easypoker.IsFlush, true)
		patternMatchingTest(t, thisHand, easypoker.IsStraightFlush, false)
		patternMatchingTest(t, thisHand, easypoker.IsRoyalFlush, false)

		h, _ := easypoker.Strand(thisHand)
		p := easypoker.HighestPokerHand(h)
		if p.Grade != easypoker.Flush {
			t.Errorf("expected %d (Flush) for %s but got %d", easypoker.Flush, thisHand, p.Grade)
		}
		queenOfHearts := easypoker.ConstructCard(french.Hearts, french.Queen)
		if p.Good.HighCard() != queenOfHearts {
			t.Errorf("expected high card of %s but got %s", queenOfHearts, p.Good.HighCard())
		}

	})

	thisHand = "🃊🃋🃍🃎🃁"
	t.Run(thisHand, func(t *testing.T) {
		patternMatchingTest(t, thisHand, easypoker.IsStraight, true)
		patternMatchingTest(t, thisHand, easypoker.IsFlush, true)
		patternMatchingTest(t, thisHand, easypoker.IsStraightFlush, true)
		patternMatchingTest(t, thisHand, easypoker.IsRoyalFlush, true)
		patternMatchingTest(t, thisHand, easypoker.HasPair, false)
		patternMatchingTest(t, thisHand, easypoker.HasTwoPair, false)

		h, _ := easypoker.Strand(thisHand)
		p := easypoker.HighestPokerHand(h)
		if p.Grade != easypoker.RoyalFlush {
			t.Errorf("expected %d (RoyalFlush) for %s but got %d", easypoker.RoyalFlush, thisHand, p.Grade)
		}

	})

	thisHand = "🂽🂭🃊🂺🂷"
	t.Run(thisHand, func(t *testing.T) {
		patternMatchingTest(t, thisHand, easypoker.HasPair, true)
		patternMatchingTest(t, thisHand, easypoker.HasTwoPair, true)
		patternMatchingTest(t, thisHand, easypoker.IsFlush, false)
		patternMatchingTest(t, thisHand, easypoker.IsRoyalFlush, false)

		h, _ := easypoker.Strand(thisHand)
		p := easypoker.HighestPokerHand(h)
		if p.Grade != easypoker.TwoPair {
			t.Errorf("expected %d (TwoPair) for %s but got %v", easypoker.TwoPair, thisHand, p)
		}
	})

	thatHand := "🃕🃂🃃🂡🃄"
	t.Run(fmt.Sprintf("%s beats %s", thisHand, thatHand), func(t *testing.T) {
		h1, _ := easypoker.Strand(thisHand)
		h2, _ := easypoker.Strand(thatHand)
		if !h1.Beats(h2) {
			t.Errorf("%s should beat %s", thisHand, thatHand)
		}
		p1 := easypoker.HighestPokerHand(h1)
		p2 := easypoker.HighestPokerHand(h2)
		if p1.Grade != easypoker.TwoPair {
			t.Errorf("expected grade %d (TwoPair) but got %d", easypoker.TwoPair, p1.Grade)
		}
		if p2.Grade != easypoker.HighCard {
			t.Errorf("expected grade %d (HighCard) but got %d", easypoker.HighCard, p2.Grade)
		}
	})

	t.Run("sort hands by poker value", func(t *testing.T) {
		inputStrings := []string{
			"🃊🃋🃍🃎🃁",
			"🂹🃙🂽🂭🃍",
			"🃙🂽🂭🃊🃎",
			"🃕🃂🃃🂡🃄",
			"🂽🂶🂷🂸🂹",
			"🂽🂭🃊🂺🂷",
			"🂶🂷🂸🂹🂺",
		}
		wantMap := map[string]easypoker.PatternGrade{
			"🃊🃋🃍🃎🃁": easypoker.RoyalFlush,
			"🂹🃙🂽🂭🃍": easypoker.FullHouse,
			"🃙🂽🂭🃊🃎": easypoker.Pair,
			"🃕🃂🃃🂡🃄": easypoker.HighCard,
			"🂽🂶🂷🂸🂹": easypoker.Flush,
			"🂽🂭🃊🂺🂷": easypoker.TwoPair,
			"🂶🂷🂸🂹🂺": easypoker.StraightFlush,
		}
		for _, thisHand := range inputStrings {
			h, _ := easypoker.Strand(thisHand)
			p := easypoker.HighestPokerHand(h)
			if p.Grade != wantMap[thisHand] {
				t.Errorf("%s should have PatternGrade %d but got %d", thisHand, wantMap[thisHand], p.Grade)
			}
		}
	})

}
