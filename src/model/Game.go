package model

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

func (game Game) GetPos(i uint, j uint) byte {
	return game.board.grid[i][j]
}

func (game *Game) CanPlay(i uint, j uint) bool {
	pos := Position{i, j}
	player := game.GetCurrentPlayer()
	return game.board.IsValidMove(player, &pos)
}

func (game *Game) Play(i uint, j uint) {
	pos := Position{i, j}
	player := game.GetCurrentPlayer()
	game.board.Play(player, &pos)
	game.turn++
}
