package model

const EMPTY byte = ' '
const BOARD_SIZE uint = 5

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

func GetOpposingColor(color byte) byte {
	if color == BLACK {
		return WHITE
	}
	return BLACK
}

func (board *Board) GetPos(i uint, j uint) byte {
	if BOARD_SIZE <= i || BOARD_SIZE <= j {
		return EMPTY
	}
	return board.grid[i][j]
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
	if 0 <= int(pos.i-1) && (pos.i+1 < BOARD_SIZE) {
		board.CheckPosition(playerColor, Position{pos.i + 1, pos.j}, chain)
		board.CheckPosition(playerColor, Position{pos.i - 1, pos.j}, chain)
	} else if BOARD_SIZE <= pos.i+1 {
		board.CheckPosition(playerColor, Position{pos.i - 1, pos.j}, chain)
	} else { // pos.i - 1 < 0
		board.CheckPosition(playerColor, Position{pos.i + 1, pos.j}, chain)
	}
	if 0 <= int(pos.j-1) && (pos.j+1 < BOARD_SIZE) {
		board.CheckPosition(playerColor, Position{pos.i, pos.j + 1}, chain)
		board.CheckPosition(playerColor, Position{pos.i, pos.j - 1}, chain)
	} else if BOARD_SIZE <= pos.j+1 {
		board.CheckPosition(playerColor, Position{pos.i, pos.j - 1}, chain)
	} else { // pos.j - 1 < 0
		board.CheckPosition(playerColor, Position{pos.i, pos.j + 1}, chain)
	}
}

func (board *Board) IsSuicide(playerColor byte, pos *Position) bool {
	chain := NewEmptyChain()
	chain.AddLiberty() // initial pos
	board.CheckNeighbours(playerColor, pos, &chain)

	// todo check special capture (9)

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
	if board.IsSuicide(player.color, pos) {
		return false
	}

	return true
}

func (board *Board) RemoveStones(chain *Chain) {
	for pos, _ := range chain.stones {
		board.grid[pos.i][pos.j] = EMPTY
	}
}

func (board *Board) CheckIfSurrounded(player *Player, pos Position) {
	if board.grid[pos.i][pos.j] == EMPTY { // it may be captured with another chain, so we can skip this
		return
	}
	chain := NewChain(pos)
	board.CheckNeighbours(GetOpposingColor(player.color), &pos, &chain)
	if chain.liberties == 0 {
		player.AddScore(chain.GetAmountOfRivalStones())
		board.RemoveStones(&chain)
	}
}

func (board *Board) Play(player *Player, pos *Position) {

	if !board.IsAValidMove(player, pos) {
		return
	}

	board.grid[pos.i][pos.j] = player.color
	player.AddMove(*pos)

	if 0 <= int(pos.i-1) && (pos.i+1) < BOARD_SIZE {
		board.CheckIfSurrounded(player, Position{pos.i + 1, pos.j})
		board.CheckIfSurrounded(player, Position{pos.i - 1, pos.j})
	} else if BOARD_SIZE <= pos.i+1 {
		board.CheckIfSurrounded(player, Position{pos.i - 1, pos.j})
	} else { // pos.i - 1 < 0
		board.CheckIfSurrounded(player, Position{pos.i + 1, pos.j})
	}
	if 0 <= int(pos.j-1) && (pos.j+1 < BOARD_SIZE) {
		board.CheckIfSurrounded(player, Position{pos.i, pos.j + 1})
		board.CheckIfSurrounded(player, Position{pos.i, pos.j - 1})
	} else if BOARD_SIZE <= pos.j+1 {
		board.CheckIfSurrounded(player, Position{pos.i, pos.j - 1})
	} else { // pos.j - 1 < 0
		board.CheckIfSurrounded(player, Position{pos.i, pos.j + 1})
	}

}
