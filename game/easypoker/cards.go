package easypoker

// Cards is a distinct datatype, because we need to compare groups of cards
// in a different way than we compare individual ones
// for example, three threes beats a high ace

type Cards []Card
