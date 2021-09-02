package model

import "math"

const EMPTY rune = ' '
const BOARD_SIZE uint = 5
const POINTS_PER_STONE uint = 1

type Position struct {
	i uint
	j uint
}

type Board struct {
	grid [BOARD_SIZE][BOARD_SIZE]rune
}

func NewBoard() Board {
	board := Board{}
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			board.grid[i][j] = EMPTY
		}
	}
	return board
}

func getOpposingColor(color rune) rune {
	if color == BLACK {
		return WHITE
	}
	return BLACK
}

func (board *Board) GetPos(i uint, j uint) rune {
	if BOARD_SIZE <= i || BOARD_SIZE <= j {
		return EMPTY
	}
	return board.grid[i][j]
}

func (board *Board) checkPosition(playerColor rune, pos Position, chain *Chain) {
	if chain.HasPos(pos) {
		return
	}
	if board.grid[pos.i][pos.j] == playerColor {
		chain.AddPos(pos)
		board.checkNeighbours(playerColor, &pos, chain)
	} else if board.grid[pos.i][pos.j] == EMPTY {
		chain.AddLiberty(pos)
	} else {
		chain.AddRival(pos)
	}

}

func (board *Board) checkNeighbours(playerColor rune, pos *Position, chain *Chain) {
	if 0 <= int(pos.i-1) && (pos.i+1 < BOARD_SIZE) {
		board.checkPosition(playerColor, Position{pos.i + 1, pos.j}, chain)
		board.checkPosition(playerColor, Position{pos.i - 1, pos.j}, chain)
	} else if BOARD_SIZE <= pos.i+1 {
		board.checkPosition(playerColor, Position{pos.i - 1, pos.j}, chain)
	} else { // pos.i - 1 < 0
		board.checkPosition(playerColor, Position{pos.i + 1, pos.j}, chain)
	}
	if 0 <= int(pos.j-1) && (pos.j+1 < BOARD_SIZE) {
		board.checkPosition(playerColor, Position{pos.i, pos.j + 1}, chain)
		board.checkPosition(playerColor, Position{pos.i, pos.j - 1}, chain)
	} else if BOARD_SIZE <= pos.j+1 {
		board.checkPosition(playerColor, Position{pos.i, pos.j - 1}, chain)
	} else { // pos.j - 1 < 0
		board.checkPosition(playerColor, Position{pos.i, pos.j + 1}, chain)
	}
}

func (board *Board) isSpecialCapture(color rune, pos *Position) bool {
	chain := NewEmptyChain()
	chain.AddLiberty(*pos) // initial pos
	board.checkNeighbours(color, pos, &chain)
	return chain.GetLiberties() <= 1
}

func (board *Board) isSuicide(playerColor rune, pos *Position) bool {
	chain := NewEmptyChain()
	chain.AddLiberty(*pos) // initial pos
	board.checkNeighbours(playerColor, pos, &chain)

	if board.isSpecialCapture(getOpposingColor(playerColor), pos) {
		return false
	}

	return chain.GetLiberties() <= 1
}

func (board *Board) IsAValidMove(player *Player, pos *Position) bool {
	if BOARD_SIZE <= pos.i || BOARD_SIZE <= pos.j {
		return false
	}
	if board.grid[pos.i][pos.j] != EMPTY {
		return false
	}
	if player.RepeatingPosition(*pos) {
		return false
	}
	if board.isSuicide(player.color, pos) {
		return false
	}

	return true
}

func (board *Board) removeStones(chain *Chain) {
	for pos, _ := range chain.stones {
		board.grid[pos.i][pos.j] = EMPTY
	}
}

func (board *Board) checkIfSurrounded(player *Player, pos Position) {
	if board.grid[pos.i][pos.j] == EMPTY { // it may be captured with another chain, so we can skip this
		return
	}
	chain := NewChain(pos)
	board.checkNeighbours(getOpposingColor(player.color), &pos, &chain)
	if !chain.HasAnyLiberties() {
		player.AddScore(chain.GetAmountOfRivalStones() * POINTS_PER_STONE)
		board.removeStones(&chain)
	}
}

func (board *Board) Play(player *Player, pos *Position) {

	if !board.IsAValidMove(player, pos) {
		return
	}

	board.grid[pos.i][pos.j] = player.color
	player.AddMove(*pos)

	if 0 <= int(pos.i-1) && (pos.i+1) < BOARD_SIZE {
		board.checkIfSurrounded(player, Position{pos.i + 1, pos.j})
		board.checkIfSurrounded(player, Position{pos.i - 1, pos.j})
	} else if BOARD_SIZE <= pos.i+1 {
		board.checkIfSurrounded(player, Position{pos.i - 1, pos.j})
	} else { // pos.i - 1 < 0
		board.checkIfSurrounded(player, Position{pos.i + 1, pos.j})
	}
	if 0 <= int(pos.j-1) && (pos.j+1 < BOARD_SIZE) {
		board.checkIfSurrounded(player, Position{pos.i, pos.j + 1})
		board.checkIfSurrounded(player, Position{pos.i, pos.j - 1})
	} else if BOARD_SIZE <= pos.j+1 {
		board.checkIfSurrounded(player, Position{pos.i, pos.j - 1})
	} else { // pos.j - 1 < 0
		board.checkIfSurrounded(player, Position{pos.i, pos.j + 1})
	}

}

func (board *Board) IsAValidDeadStone(i uint, j uint) bool {
	return board.grid[i][j] != EMPTY
}

func (board *Board) MarkAsDeadStone(i uint, j uint, blackPlayer *Player, whitePlayer *Player) {
	if !board.IsAValidDeadStone(i, j) {
		return
	}
	if board.grid[i][j] == BLACK {
		whitePlayer.AddScore(POINTS_PER_STONE)
	} else {
		blackPlayer.AddScore(POINTS_PER_STONE)
	}
	board.grid[i][j] = getOpposingColor(board.grid[i][j])
}

func (board *Board) paintWithClosestColoredStone(i uint, j uint) { // todo ref
	var x, y uint
	var minI, minJ = float64(BOARD_SIZE), float64(BOARD_SIZE)
	for x = 0; x < BOARD_SIZE; x++ {
		absValueI := math.Abs(float64(i) - float64(x))
		for y = 0; y < BOARD_SIZE; y++ {
			absValueJ := math.Abs(float64(j) - float64(y))
			if 1 <= (absValueI+absValueJ) && (absValueI+absValueJ) <= (minI+minJ) {
				board.grid[i][j] = board.grid[x][y]
				minI = absValueI
				minJ = absValueJ
			}
		}
	}
}

func (board *Board) AddTerritoryPoints(blackPlayer *Player, whitePlayer *Player) {

	var i, j uint
	for i = 0; i < BOARD_SIZE; i++ { // todo, refactor to a map saving the empty stones
		for j = 0; j < BOARD_SIZE; j++ {
			if board.grid[i][j] == EMPTY {
				board.paintWithClosestColoredStone(i, j)
				if board.grid[i][j] == BLACK {
					blackPlayer.AddScore(POINTS_PER_STONE)
				} else {
					whitePlayer.AddScore(POINTS_PER_STONE)
				}
			}
		}
	}

}
