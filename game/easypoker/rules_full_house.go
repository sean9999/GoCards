package easypoker

var IsFullHouse PokerPattern = PokerPattern{
	Grade:       FullHouse,
	Declaration: "is full-house",
	Description: "hand is composed of a pair and a three of a kind",
	Func:        fn_IsFullHouse,
}

var fn_IsFullHouse PatternFunc = func(cards Cards) (bool, PokerHand) {
	hasTwins, twinHand := HasPair.Func(cards)
	hasTriplets, trebleHand := HasThreeOfAKind.Func(twinHand.Remaining)
	switch {
	case !hasTwins && !hasTriplets:
		return false, ZeroPokerHand
	case !hasTwins && hasTriplets:
		return false, trebleHand
	case hasTwins && !hasTriplets:
		return false, twinHand
	default:
		//	happy path
		ph := PokerHand{
			Grade:     FullHouse,
			Good:      append(twinHand.Good, trebleHand.Good...),
			Remaining: []Card{},
		}
		return true, ph
	}
}
