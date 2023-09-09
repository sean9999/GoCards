package easypoker

var GetHighCard = PokerPattern{
	Declaration: "has a high card",
	Description: "get the highest ranking card",
	Grade:       HighCard,
	Func:        fn_GetHighCard,
}

var fn_GetHighCard PatternFunc = func(cards Cards) (bool, PokerHand) {
	cards.SortLowToHigh()
	chop := len(cards) - 2
	head := cards[:chop]
	tail := cards[chop:]
	ph := PokerHand{
		Good:      tail,
		Remaining: head,
		Grade:     HighCard,
	}
	return true, ph
}
