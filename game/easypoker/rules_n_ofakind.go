package easypoker

var fn_HasFourOfAKind PatternFunc = OfAKindFactory(4)
var fn_HasThreeOfAKind PatternFunc = OfAKindFactory(3)
var fn_HasPair PatternFunc = OfAKindFactory(2)

var HasPair = PokerPattern{
	Grade:       Pair,
	Declaration: "has a pair",
	Description: "two cards within have the same rank",
	Func:        fn_HasPair,
}

var HasThreeOfAKind = PokerPattern{
	Grade:       ThreeOfAKind,
	Declaration: "has three of a kind",
	Description: "three cards within have the same rank",
	Func:        fn_HasThreeOfAKind,
}

var HasFourOfAKind = PokerPattern{
	Grade:       FourOfAKind,
	Description: "four cards within have the same rank",
	Declaration: "has four of a kind",

	Func: fn_HasFourOfAKind,
}
