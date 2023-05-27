/*
package easypoker represents a game of poker that is easy to reason about and trivial to use in tests and PoCs.

It has the following features:

- uses a standard 52-card french-suited deck. The deck most of us are used to
- no jokers
- no betting
- no wild cards
- no extra draws
- one hand of 5 cards
- one hand is dealt to each player. That's it. Whoever wins, wins.
- a Game is composed of as many rounds as it takes to exhaust the deck
- Therefore, a game cannot be played where more than 10 people sign up to join
- You pass in your random seed at the beginning, allowing for deterministic games, if desired
*/
package easypoker
