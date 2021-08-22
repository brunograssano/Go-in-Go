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

func (game Game) GetCurrentPlayer() *Player {
	var player *Player
	if game.turn%2 == 0 {
		player = &game.blackPlayer
	} else {
		player = &game.whitePlayer
	}
	return player
}

func (game *Game) CanPlay(pos *Position) bool {
	player := game.GetCurrentPlayer()
	return game.board.IsValidMove(player, pos)
}

func (game *Game) Play(pos *Position) {
	player := game.GetCurrentPlayer()
	game.board.Play(player, pos)
	game.turn++
}
