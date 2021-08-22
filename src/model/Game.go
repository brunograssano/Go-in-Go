package main

const BLACK = 'B'
const WHITE = 'W'

type Game struct {
	blackPlayer Player
	whitePlayer Player
	board       Board
	turn        uint // todo wip
}

func NewGame() Game {
	game := Game{NewPlayer(BLACK), NewPlayer(WHITE), NewBoard(), 0}
	return game
}

func (game *Game) Play(pos *Position) {
	var player *Player
	if game.turn%2 == 0 {
		player = &game.blackPlayer
	} else {
		player = &game.whitePlayer
	}
	game.board.Play(player, pos)
	game.turn++
}
