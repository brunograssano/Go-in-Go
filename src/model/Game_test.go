package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTheGameStartsEmpty01(t *testing.T) {
	assert := assert.New(t)
	game := NewGame()
	var i, j uint
	for i = 0; i < BOARD_SIZE; i++ {
		for j = 0; j < BOARD_SIZE; j++ {
			assert.Equal(EMPTY, game.GetPos(i, j), "The game is empty.")
		}
	}
}

/*
	1)		2)		3)		4)		5)
  	0 1 .	0 1 .	0 1 .	0 1 .	0 1 .
0  			B		B W		B W		  W
1							W		W
.
*/
func TestBlackIsSurroundedByWhiteOnTheBorders02(t *testing.T) {
	assert := assert.New(t)
	game := NewGame()

	game.Play(0, 0) // Black
	assert.Equal(BLACK, game.GetPos(0, 0), "Black stone in (0,0)")
	game.Play(0, 1) // White
	assert.Equal(WHITE, game.GetPos(0, 1), "White stone in (0,1)")
	game.Play(4, 4) // Black
	assert.Equal(BLACK, game.GetPos(4, 4), "Black stone in (4,4)")
	game.Play(1, 0) // White, captures Black
	assert.Equal(WHITE, game.GetPos(1, 0), "White stone in (1,0)")
	assert.Equal(WHITE, game.GetPos(1, 0), "White stone still in (0,1)")
	assert.Equal(EMPTY, game.GetPos(0, 0), "Black stone in (0,0) was captured")
	assert.Equal(BLACK, game.GetPos(4, 4), "Black stone still in (4,4)")
}

/*
	1)		2)		3)		4)		5)		6)
  	0 1 2	0 1 2	0 1 2	0 1 2	0 1 2	0 1 2
0  	  		  B		W B		W B W	W B W	W   W
1									  W		  W
.
*/
func TestOneBlackStoneIsSurroundedByThreeWhiteStones03(t *testing.T) {
	assert := assert.New(t)
	game := NewGame()

	game.Play(0, 1) // Black
	game.Play(0, 0) // White
	game.PassTurn() // Black
	game.Play(0, 2) // White
	assert.Equal(BLACK, game.GetPos(0, 1), "Black stone in (0,1)")
	game.PassTurn() // Black
	game.Play(1, 1) // White
	assert.Equal(WHITE, game.GetPos(0, 0), "White stone in (0,0)")
	assert.Equal(WHITE, game.GetPos(1, 1), "White stone in (1,1)")
	assert.Equal(EMPTY, game.GetPos(0, 1), "Black stone in (0,1) was captured")
	assert.Equal(WHITE, game.GetPos(0, 2), "White stone in (0,2)")
}

/*
  0 1 2 3
0   W W
1 W B B W
2   W W
*/
func TestTwoBlackStonesAreSurroundedByWhiteStones04(t *testing.T) {
	assert := assert.New(t)
	game := NewGame()

	game.Play(1, 1) // Black
	game.Play(0, 1) // White
	game.Play(1, 2) // Black
	game.Play(0, 2) // White
	game.PassTurn() // Black
	game.Play(1, 0) // White
	game.PassTurn() // Black
	game.Play(1, 3) // White
	game.PassTurn() // Black
	game.Play(2, 1) // White
	game.PassTurn() // Black

	assert.Equal(BLACK, game.GetPos(1, 1), "There is a black stone in (1,1)")
	assert.Equal(BLACK, game.GetPos(1, 2), "There is a black stone in (1,2)")

	game.Play(2, 2) // White

	assert.Equal(EMPTY, game.GetPos(1, 1), "The black stone in (1,1) was captured")
	assert.Equal(EMPTY, game.GetPos(1, 2), "The black stone in (1,2) was captured")
}
