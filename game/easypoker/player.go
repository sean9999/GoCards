package easypoker

type Player struct {
	Name string
}

func NewPlayer(Name string) Player {
	return Player{
		Name,
	}
}
