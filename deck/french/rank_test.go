package french_test

import (
	"testing"

	"github.com/sean9999/GoCards/deck/french"
	"github.com/stretchr/testify/assert"
)

func TestRank(t *testing.T) {
	assert.Equal(t, "Ace", french.Ace.String())
	assert.Equal(t, "Two", french.Two.String())
	assert.Equal(t, "Joker", french.Joker.String())
	assert.Equal(t, french.Rank(0), french.ZeroRank)
}
