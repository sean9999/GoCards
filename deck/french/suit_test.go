package french_test

import (
	"testing"

	"github.com/sean9999/GoCards/deck/french"
	"github.com/stretchr/testify/assert"
)

func TestSuit(t *testing.T) {

	t.Run("zero suit", func(t *testing.T) {

		var noSuit french.Suit

		t.Run("zero suit is not a valid suit", func(t *testing.T) {
			isValid, err := noSuit.Validate()
			assert.False(t, isValid)
			assert.ErrorIs(t, err, french.ErrOutofRange)
		})

		t.Run("zero suit range produces a zero range", func(t *testing.T) {
			zeroRange := french.SuitRange{}
			gotRange := noSuit.Range()
			assert.Equal(t, zeroRange, gotRange)
		})

	})

	t.Run("real suits", func(t *testing.T) {
		for thisSuit, thisRange := range french.LegalSuitRanges {
			switch thisSuit {
			case french.Hearts, french.Clubs, french.Spades, french.Diamonds:
				assert.Equal(t, french.Ace, thisRange.LowerBound.Rank(), "lowest rank should be Ace")
				assert.Equal(t, french.King, thisRange.UpperBound.Rank(), "highest rank should be King")

			case french.Black, french.Red, french.White:
				assert.Equal(t, thisRange.LowerBound, thisRange.UpperBound, "suit should contain only one card")
				assert.Equal(t, french.Joker, thisRange.LowerBound.Rank(), "that card should be Joker")
			}
		}
	})

	t.Run("what beats what", func(t *testing.T) {
		assert.False(t, french.Clubs.Beats(french.Diamonds), "diamonds beats clubs")
		assert.True(t, french.Hearts.Beats(french.Clubs), "hearts beats clubs")
		assert.True(t, french.Spades.Beats(french.Hearts), "spades beats hearts")
		assert.True(t, french.Spades.Beats(french.Clubs), "spades beats clubs")
		assert.True(t, french.Hearts.Beats(french.Diamonds), "hearts beats diamonds")
		assert.True(t, french.Spades.Beats(french.Diamonds), "spades beats diamonds")
	})

}
