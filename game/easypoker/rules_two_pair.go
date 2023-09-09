package easypoker

var HasTwoPair = PokerPattern{
	Declaration: "contains two pairs",
	Description: "there are two pairs in the hand",
	Grade:       TwoPair,
	Func:        fn_HasTwoPair,
}

var fn_HasTwoPair PatternFunc = func(cards Cards) (bool, PokerHand) {
	ph := ZeroPokerHand
	hasFirstPair, firstPokerHand := HasPair.Func(cards)
	if !hasFirstPair {
		return false, ph
	}
	hasSecondPair, secondPokerHand := HasPair.Func(firstPokerHand.Remaining)
	if !hasSecondPair {
		//	as a consolation, PokerHand contains the one pair we did match on
		return false, firstPokerHand
	}
	//	happy path
	ph.Remaining = secondPokerHand.Remaining
	ph.Good = append(firstPokerHand.Good, secondPokerHand.Good...).AsLowToHigh()
	ph.Grade = TwoPair
	return true, ph
}
