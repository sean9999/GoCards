package french_test

import (
	"testing"

	"github.com/sean9999/GoCards/deck/french"
	"github.com/stretchr/testify/assert"
)

func TestCard(t *testing.T) {
	c := french.MakeCard(french.Ace, french.Spades)

	isValid, err := c.Validate()
	if err != nil {
		t.Fatal(err)
	}
	assert.True(t, isValid)

	assert.Equal(t, c.Rank(), french.Ace)
	suit, err := c.Suit()
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, suit, french.Spades)
	assert.Equal(t, c.FallsWithin(), french.SpadesRange)
	assert.Equal(t, "🂡", c.String())
	assert.Equal(t, "Ace of Spades", c.Word())
}
