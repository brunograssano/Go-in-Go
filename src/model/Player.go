package main

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
