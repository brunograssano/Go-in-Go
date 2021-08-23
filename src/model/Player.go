package model

type Player struct {
	score         uint
	moves         uint
	color         byte
	previousMoves map[Position]Position
}

func NewPlayer(color byte) Player {
	player := Player{0, 0, color, make(map[Position]Position)}
	return player
}

func (player *Player) AddScore(points uint) {
	player.score += points
}

func (player *Player) RepeatingPosition(pos Position) bool {
	pos, played := player.previousMoves[pos]
	if played {
		return true
	}
	return false
}

func (player *Player) AddMove(pos Position) {
	player.previousMoves[pos] = pos
}
