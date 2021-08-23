package model

const WHITE_EXTRA_POINTS uint = 7

type Player struct {
	score         uint
	moves         uint
	color         byte
	previousMoves map[Position]Position
}

func NewPlayer(color byte) Player {
	var initialScore uint = 0
	if color == WHITE {
		initialScore = WHITE_EXTRA_POINTS
	}
	player := Player{initialScore, 0, color, make(map[Position]Position)}
	return player
}

func (player *Player) AddScore(points uint) {
	player.score += points
}

func (player *Player) RepeatingPosition(pos Position) bool {
	pos, played := player.previousMoves[pos]
	return played
}

func (player *Player) AddMove(pos Position) {
	player.previousMoves[pos] = pos
}
