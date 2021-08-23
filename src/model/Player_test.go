package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestANewBlackPlayerHasZeroPoints01(t *testing.T) {
	player := NewPlayer(BLACK)
	assert.Equal(t, uint(0), player.score, "The score of a new black player should be 0.")
}

func TestAddingFivePointsToANewBlackPlayerResultsInFivePoints02(t *testing.T) {
	assert := assert.New(t)
	player := NewPlayer(BLACK)
	player.AddScore(1)
	assert.Equal(uint(1), player.score, "The score should be 1.")
	player.AddScore(2)
	assert.Equal(uint(3), player.score, "The score should be 3.")
	player.AddScore(2)
	assert.Equal(uint(5), player.score, "The score should be 5.")
}

func TestAWhitePlayerStartsWithSevenPoints03(t *testing.T) {
	player := NewPlayer(WHITE)
	assert.Equal(t, WHITE_EXTRA_POINTS, player.score, "The score of a new white player should be %d.", WHITE_EXTRA_POINTS)
}

func TestAPlayerCannotRepeatMovesPreviouslyPlayed04(t *testing.T) {
	assert := assert.New(t)
	pos := Position{0, 0}
	pos2 := Position{2, 2}
	pos3 := Position{3, 3}
	player := NewPlayer(WHITE)
	assert.False(player.RepeatingPosition(pos), "The player does not have the position (0,0)")
	assert.False(player.RepeatingPosition(pos2), "The player does not have the position (2,2)")
	assert.False(player.RepeatingPosition(pos3), "The player does not have the position (3,3)")
	player.AddMove(pos)
	player.AddMove(pos2)
	player.AddMove(pos3)
	assert.True(player.RepeatingPosition(pos), "The player is repeating the position (0,0)")
	assert.True(player.RepeatingPosition(pos2), "The player is repeating the position (2,2)")
	assert.True(player.RepeatingPosition(pos3), "The player is repeating the position (3,3)")
}
