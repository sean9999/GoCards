package french_test

import (
	"fmt"
	"testing"

	"github.com/sean9999/GoCards/deck/french"
)

func assertScalarEquals[C comparable](t *testing.T, got, want C) {
	t.Helper()
	if want != got {
		t.Errorf("want %v but got %v", want, got)
	}
}

func cardsAreEqual(t *testing.T, got, want []french.Card) (bool, error) {
	t.Helper()
	if len(got) != len(want) {
		return false, fmt.Errorf("got %v but wanted %v", got, want)
	}
	for i := 0; i < len(got); i++ {
		if got[i] != want[i] {
			return false, fmt.Errorf("got %s but wanted %s", french.Cards(got).Strand(), french.Cards(want).Strand())
		}
	}
	return true, nil
}

func GenerateHand(str string) french.Cards {
	chars := []rune(str)
	cards := make([]french.Card, 0, len(chars))
	for _, char := range chars {
		c, _ := french.CardFromRune(char)
		cards = append(cards, c)
	}
	return cards
}

func TestNewDeck(t *testing.T) {

	t.Run("New unshuffled deck", func(t *testing.T) {

		uStock := french.NewDeck().DealOut()
		t.Run("has 54 cards", func(t *testing.T) {
			assertScalarEquals[int](t, len(uStock), 54)
		})

		wantedHandStrings := []string{
			"🃑🃒🃓🃔🃕",
			"🃖🃗🃘🃙🃚",
			"🃛🃝🃞🃁🃂",
			"🃃🃄🃅🃆🃇",
			"🃈🃉🃊🃋🃍",
			"🃎🂱🂲🂳🂴",
			"🂵🂶🂷🂸🂹",
			"🂺🂻🂽🂾🂡",
			"🂢🂣🂤🂥🂦",
			"🂧🂨🂩🂪🂫",
		}

		for handNumber, wantString := range wantedHandStrings {
			t.Run(fmt.Sprintf("hand %d should be %s", handNumber, wantString), func(t *testing.T) {
				want := GenerateHand(wantString)
				got, err := uStock.Draw(5)
				if err != nil {
					t.Error(err)
				} else {
					equal, err := cardsAreEqual(t, got, want)
					if !equal {
						t.Error(err)
					}
				}
			})
		}

		t.Run("4 cards left", func(t *testing.T) {
			assertScalarEquals[int](t, len(uStock), 4)
		})

	})

}
