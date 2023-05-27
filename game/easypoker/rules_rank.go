package easypoker

func (r Rank) Beats(q Rank) bool {
	return r > q
}
