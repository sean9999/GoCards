package easypoker

var IsStraightFlush = PokerPattern{
	Declaration: "is a straight flush",
	Description: "cards are in sequence and of the same suit",
	Grade:       StraightFlush,
	Func:        fn_IsStraightFlush,
}

var fn_IsStraightFlush = func(cards Cards) (bool, PokerHand) {
	ok := false
	ok, flushPokerHand := IsFlush.Func(cards)
	if ok {
		ok, _ := IsStraight.Func(cards)
		if ok {
			return true, flushPokerHand
		}
	}
	return false, ZeroPokerHand
}
