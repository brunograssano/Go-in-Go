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
  0 1
0 B W
1 W
*/
func TestBlackIsSurroundedByWhiteOnTheUpperLeftBorder02(t *testing.T) {
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
  0 1 2
0 W B W
1   W
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

/*
  0 1 2 3
0   B B
1 B W W B
2   B W B
3     B
*/
func TestThreeWhiteStonesAreSurroundedByBlackStones05(t *testing.T) {
	assert := assert.New(t)
	game := NewGame()

	game.Play(0, 1) // Black
	game.Play(1, 1) // White
	game.Play(0, 2) // Black
	game.Play(1, 2) // White
	game.Play(1, 0) // Black
	game.Play(2, 2) // White
	game.Play(3, 2) // Black
	game.PassTurn() // White
	game.Play(2, 1) // Black
	game.PassTurn() // White
	game.Play(2, 3) // Black
	game.PassTurn() // White

	assert.Equal(WHITE, game.GetPos(1, 1), "There is a white stone in (1,1)")
	assert.Equal(WHITE, game.GetPos(1, 2), "There is a white stone in (1,2)")
	assert.Equal(WHITE, game.GetPos(2, 2), "There is a white stone in (2,2)")

	game.Play(1, 3) // Black

	assert.Equal(EMPTY, game.GetPos(1, 1), "The white stone in (1,1) was captured")
	assert.Equal(EMPTY, game.GetPos(1, 2), "The white stone in (1,2) was captured")
	assert.Equal(EMPTY, game.GetPos(2, 2), "The white stone in (2,2) was captured")
}

/* White tries to play (2,2)
  0 1 2 3
0   B B
1 B W W B
2   B   B
3     B
*/
func TestYouCantCaptureYourselfIfYouAreSurrounded06(t *testing.T) {
	assert := assert.New(t)
	game := NewGame()

	game.Play(0, 1) // Black
	game.Play(1, 1) // White
	game.Play(0, 2) // Black
	game.Play(1, 2) // White
	game.Play(1, 0) // Black
	game.PassTurn() // White
	game.Play(3, 2) // Black
	game.PassTurn() // White
	game.Play(2, 1) // Black
	game.PassTurn() // White
	game.Play(2, 3) // Black
	game.PassTurn() // White
	game.Play(1, 3) // Black

	assert.Equal(EMPTY, game.GetPos(2, 2), "(2,2) is empty")
	assert.False(game.CanPlay(2, 2), "(2,2) is not a valid move for white")
	game.Play(2, 2) // White

	assert.Equal(EMPTY, game.GetPos(2, 2), "(2,2) is still empty after playing it")
}

/*
  0 1 2 3 4
0 W X B W
1 B B B B W
2 W B W W
3   W
*/
func TestSpecialCapture07(t *testing.T) {
	assert := assert.New(t)
	game := NewGame()

	game.Play(0, 2) // Black
	game.Play(0, 0) // White
	game.Play(1, 0) // Black
	game.Play(0, 3) // White
	game.Play(1, 1) // Black
	game.Play(2, 0) // White
	game.Play(1, 2) // Black
	game.Play(3, 1) // White
	game.Play(2, 1) // Black
	game.Play(2, 2) // White
	game.Play(1, 3) // Black
	game.Play(1, 4) // White
	game.Play(1, 3) // Black
	game.Play(2, 3) // White
	game.PassTurn() // Black

	assert.Equal(EMPTY, game.GetPos(0, 1), "(0,1) is empty")
	assert.True(game.CanPlay(0, 1), "(0,1) is a valid move for white")
	game.Play(0, 1) // White

	assert.Equal(WHITE, game.GetPos(0, 1), "(0,1) belongs to White")
	assert.Equal(WHITE, game.GetPos(0, 0), "(0,0) belongs to White")
	assert.Equal(EMPTY, game.GetPos(0, 2), "(0,2) is empty")
	assert.Equal(EMPTY, game.GetPos(1, 1), "(1,1) is empty")
	assert.Equal(EMPTY, game.GetPos(1, 0), "(1,0) is empty")
	assert.Equal(EMPTY, game.GetPos(2, 1), "(2,1) is empty")
}

/*
  0 1 2 3 4
0 b b b B W
1 B B B B W
2 W W W W w
3 w w w w w
4 w w w w w
*/
func TestBlackCapturesThreeTerritoriesWhileWhiteCapturesElevenAtTheEndOfTheGame(t *testing.T) {
	assert := assert.New(t)
	game := NewGame()

	game.Play(1, 0) // Black
	game.Play(2, 0) // White
	game.Play(1, 1) // Black
	game.Play(2, 1) // White
	game.Play(1, 2) // Black
	game.Play(2, 2) // White
	game.Play(1, 3) // Black
	game.Play(2, 3) // White
	game.Play(0, 3) // Black
	game.Play(0, 4) // White
	game.PassTurn() // Black
	game.Play(1, 4) // White

	blackScore, whiteScore := game.GetScore()

	assert.Equal(uint(3), blackScore, "Black has 3 points")
	assert.Equal(11+WHITE_EXTRA_POINTS, whiteScore, "White has %d points", 11+WHITE_EXTRA_POINTS)

}

/*
  0 1 2 3 4
0 b b b b b
1 B B B B B
2 W W W W W
3 w w w w w
4 w w w w w
*/
func TestBlackCapturesFiveTerritoriesWhileWhiteCapturesTenAtTheEndOfTheGame(t *testing.T) {
	assert := assert.New(t)
	game := NewGame()

	game.Play(1, 0) // Black
	game.Play(2, 0) // White
	game.Play(1, 1) // Black
	game.Play(2, 1) // White
	game.Play(1, 2) // Black
	game.Play(2, 2) // White
	game.Play(1, 3) // Black
	game.Play(2, 3) // White
	game.Play(1, 4) // Black
	game.Play(2, 4) // White

	blackScore, whiteScore := game.GetScore()

	assert.Equal(uint(5), blackScore, "Black has 5 points")
	assert.Equal(10+WHITE_EXTRA_POINTS, whiteScore, "White has %d points", 10+WHITE_EXTRA_POINTS)

}
