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

func (board *Board) CheckPosition(playerColor byte, pos Position, chain *Chain) {
	if chain.HasPos(pos) {
		return
	}
	if board.grid[pos.i][pos.j] == playerColor {
		chain.AddPos(pos)
		board.CheckNeighbours(playerColor, &pos, chain)
	} else if board.grid[pos.i][pos.j] == EMPTY {
		chain.AddLiberty()
	} else {
		chain.AddRival()
	}

}

func (board *Board) CheckNeighbours(playerColor byte, pos *Position, chain *Chain) {
	if 0 <= pos.i-1 && pos.i+1 <= BOARD_SIZE { // todo refactor
		board.CheckPosition(playerColor, Position{pos.i + 1, pos.j}, chain)
		board.CheckPosition(playerColor, Position{pos.i - 1, pos.j}, chain)
	} else if BOARD_SIZE <= pos.i+1 {
		board.CheckPosition(playerColor, Position{pos.i - 1, pos.j}, chain)
	} else { // pos.i - 1 < 0
		board.CheckPosition(playerColor, Position{pos.i + 1, pos.j}, chain)
	}
	if 0 <= pos.j-1 && pos.j+1 <= BOARD_SIZE {
		board.CheckPosition(playerColor, Position{pos.i, pos.j + 1}, chain)
		board.CheckPosition(playerColor, Position{pos.i, pos.j - 1}, chain)
	} else if BOARD_SIZE <= pos.j+1 {
		board.CheckPosition(playerColor, Position{pos.i, pos.j - 1}, chain)
	} else { // pos.j - 1 < 0
		board.CheckPosition(playerColor, Position{pos.i, pos.j + 1}, chain)
	}
}

func (board *Board) IsSuicide(playerColor byte, pos *Position) bool {
	chain := NewChain(*pos)
	board.CheckNeighbours(playerColor, pos, &chain)
	// todo check special capture (9)

	return chain.liberties <= 1
}

func (board *Board) IsValidMove(player *Player, pos *Position) bool {
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

	if !board.IsValidMove(player, pos) {
		return
	}

	board.grid[pos.i][pos.j] = player.color
	player.AddMove(*pos)

	// todo check capturing

}
