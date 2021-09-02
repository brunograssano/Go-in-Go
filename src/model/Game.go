package model

const BLACK rune = 'B'
const WHITE rune = 'W'

type Game struct {
	blackPlayer              Player
	whitePlayer              Player
	board                    Board
	turn                     uint
	amountOfContinuousPasses uint
}

func NewGame() Game {
	game := Game{NewPlayer(BLACK), NewPlayer(WHITE), NewBoard(), 0, 0}
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

func (game *Game) IsGameOver() bool {
	return game.amountOfContinuousPasses > 1
}

func (game *Game) IsAValidDeadStone(i uint, j uint) bool {
	return game.board.IsAValidDeadStone(i, j)
}

func (game *Game) MarkAsDeadStone(i uint, j uint) {
	game.board.MarkAsDeadStone(i, j, &game.blackPlayer, &game.whitePlayer)
}

func (game *Game) GetScore() (uint, uint) {
	game.board.AddTerritoryPoints(&game.blackPlayer, &game.whitePlayer)
	return game.blackPlayer.GetScore(), game.whitePlayer.GetScore()
}

func (game *Game) PassTurn() {
	game.turn++
	game.amountOfContinuousPasses++
}

func (game *Game) Play(i uint, j uint) {
	pos := Position{i, j}
	player := game.GetCurrentPlayer()
	game.board.Play(player, &pos)
	game.turn++
	game.amountOfContinuousPasses = 0
}
