package main

const EMPTY = ' '
const BOARD_SIZE = 5

type Position struct {
	i uint
	j uint
}

type Board struct {
	grid [BOARD_SIZE][BOARD_SIZE]byte
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

func (board *Board) GetChainAux(playerColor byte, pos Position, chain map[Position]Position) { // maps are reference types
	_, visited := chain[pos]
	if visited {
		return
	}
	if board.grid[pos.i][pos.j] == playerColor {
		chain[pos] = pos
		board.CheckNeighbours(playerColor, &pos, chain)
	}
}

func (board *Board) CheckNeighbours(playerColor byte, pos *Position, chain map[Position]Position) {
	var nextPos Position
	if 0 <= pos.i-1 && pos.i+1 <= BOARD_SIZE { // todo refactor
		nextPos = Position{pos.i + 1, pos.j}
		board.GetChainAux(playerColor, nextPos, chain)
		nextPos = Position{pos.i - 1, pos.j}
		board.GetChainAux(playerColor, nextPos, chain)
	} else if BOARD_SIZE <= pos.i+1 {
		nextPos = Position{pos.i - 1, pos.j}
		board.GetChainAux(playerColor, nextPos, chain)
	} else { // pos.i - 1 < 0
		nextPos = Position{pos.i + 1, pos.j}
		board.GetChainAux(playerColor, nextPos, chain)
	}
	if 0 <= pos.j-1 && pos.j+1 <= BOARD_SIZE {
		nextPos = Position{pos.i, pos.j + 1}
		board.GetChainAux(playerColor, nextPos, chain)
		nextPos = Position{pos.i, pos.j - 1}
		board.GetChainAux(playerColor, nextPos, chain)
	} else if BOARD_SIZE <= pos.j+1 {
		nextPos = Position{pos.i, pos.j - 1}
		board.GetChainAux(playerColor, nextPos, chain)
	} else { // pos.j - 1 < 0
		nextPos = Position{pos.i, pos.j + 1}
		board.GetChainAux(playerColor, nextPos, chain)
	}
}

func (board *Board) GetChain(playerColor byte, pos *Position) map[Position]Position {

	chain := make(map[Position]Position)
	chain[*pos] = *pos
	board.CheckNeighbours(playerColor, pos, chain)

	return chain
}

func (board *Board) AddIfLiberty(pos Position, liberties *uint) {
	if board.grid[pos.i][pos.j] == EMPTY {
		*liberties++
	}
}

func (board *Board) HasMoreThanOneLiberties(chain map[Position]Position) bool {
	var liberties uint = 0
	var nextPos Position
	for pos := range chain {
		if 0 <= pos.i-1 && pos.i+1 <= BOARD_SIZE { // todo refactor
			nextPos = Position{pos.i + 1, pos.j}
			board.AddIfLiberty(nextPos, &liberties)
			nextPos = Position{pos.i - 1, pos.j}
			board.AddIfLiberty(nextPos, &liberties)
		} else if BOARD_SIZE <= pos.i+1 {
			nextPos = Position{pos.i - 1, pos.j}
			board.AddIfLiberty(nextPos, &liberties)
		} else { // pos.i - 1 < 0
			nextPos = Position{pos.i + 1, pos.j}
			board.AddIfLiberty(nextPos, &liberties)
		}
		if 0 <= pos.j-1 && pos.j+1 <= BOARD_SIZE {
			nextPos = Position{pos.i, pos.j + 1}
			board.AddIfLiberty(nextPos, &liberties)
			nextPos = Position{pos.i, pos.j - 1}
			board.AddIfLiberty(nextPos, &liberties)
		} else if BOARD_SIZE <= pos.j+1 {
			nextPos = Position{pos.i, pos.j - 1}
			board.AddIfLiberty(nextPos, &liberties)
		} else { // pos.j - 1 < 0
			nextPos = Position{pos.i, pos.j + 1}
			board.AddIfLiberty(nextPos, &liberties)
		}
		if liberties > 1 {
			return true
		}
	}
	return false
}

func (board *Board) IsSuicide(playerColor byte, pos *Position) bool {
	var rivalPlayerColor byte
	if playerColor == BLACK {
		rivalPlayerColor = WHITE
	} else {
		rivalPlayerColor = BLACK
	}
	chain := board.GetChain(playerColor, pos)

	// todo check special capture (9)

	return board.HasMoreThanOneLiberties(chain)
}

func (board *Board) ValidPlay(player *Player, pos *Position) bool {
	if BOARD_SIZE <= pos.i || BOARD_SIZE <= pos.j {
		return false
	}
	if board.grid[pos.i][pos.j] != EMPTY {
		return false
	}
	if player.RepeatingPosition(*pos) {
		return false
	}
	if board.IsSuicide(player.color, pos) {
		return false
	}

	return true
}

func (board *Board) Play(player *Player, pos *Position) {

	if !board.ValidPlay(player, pos) {
		return
	}

	board.grid[pos.i][pos.j] = player.color
	player.AddMove(*pos)

	// todo check capturing

}
