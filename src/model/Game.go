package model

const BLACK rune = 'B'
const WHITE rune = 'W'

type Game struct {
	blackPlayer Player
	whitePlayer Player
	board       Board
	turn        uint
}

func NewGame() Game {
	game := Game{NewPlayer(BLACK), NewPlayer(WHITE), NewBoard(), 0}
	return game
}

func (game Game) GetCurrentPlayer() *Player {
	if game.turn%2 == 0 {
		return &game.blackPlayer
	}
	return &game.whitePlayer
}

func (game Game) GetPos(i uint, j uint) rune {
	return game.board.GetPos(i, j)
}

func (game *Game) CanPlay(i uint, j uint) bool {
	pos := Position{i, j}
	player := game.GetCurrentPlayer()
	return game.board.IsAValidMove(player, &pos)
}

func (game *Game) PassTurn() {
	game.turn++
}

func (game *Game) Play(i uint, j uint) {
	pos := Position{i, j}
	player := game.GetCurrentPlayer()
	game.board.Play(player, &pos)
	game.turn++
}
