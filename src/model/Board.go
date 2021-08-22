package main

const EMPTY = ' '

type Position struct {
	i uint
	j uint
}

type Board struct {
	grid [5][5]byte
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

func (board *Board) Play(player *Player, pos *Position) {

	// todo check if valid pos & move

	board.grid[pos.i][pos.j] = player.color

	// todo check capturing

}
