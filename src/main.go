package main

import (
	"./model"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const EXIT byte = 'E'

func PrintBoard(game *model.Game) {
	var i, j uint
	fmt.Printf("   ")
	for i = 0; i < model.BOARD_SIZE; i++ {
		fmt.Printf("  %d ", i)
	}
	fmt.Printf("\n")

	for i = 0; i < model.BOARD_SIZE; i++ {
		fmt.Printf(" %d ", i)
		for j = 0; j < model.BOARD_SIZE; j++ {
			fmt.Printf("| %c ", game.GetPos(i, j))
		}
		fmt.Printf("| \n")
	}
}

func ReadCharacter(reader *bufio.Reader, charID string) uint {
	fmt.Printf(charID)
	char, _, err := reader.ReadRune()
	if err != nil {
		fmt.Println(err)
	}
	_, err = reader.Discard(reader.Buffered())
	if err != nil {
		fmt.Println(err)
	}
	x, err := strconv.Atoi(string(char))
	if err != nil {
		fmt.Println(err)
		x = -1
	}
	return uint(x)
}

func ReadInput(game *model.Game) (uint, uint) {
	reader := bufio.NewReader(os.Stdin)
	var validMove = false
	var i, j uint
	for !validMove {
		i = ReadCharacter(reader, "i:")
		j = ReadCharacter(reader, "j:")
		validMove = true
		if !game.CanPlay(i, j) {
			fmt.Printf("Not a valid move! Please insert for 'i' and 'j' a value between 0 and %d.\n", model.BOARD_SIZE-1)
			fmt.Printf("Remember that you can only add a 'stone' if the place is empty, that you can't repeat previous moves, " +
				"and that you can't capture yourself!\n")
			validMove = false
		}
	}
	return i, j
}

func PrintInfo(game *model.Game) {
	player := game.GetCurrentPlayer()
	turnMessage := "White's turn\n"
	if player.GetColor() == model.BLACK {
		turnMessage = "Black's turn\n"
	}
	fmt.Printf(turnMessage)
}

func GameLoop() {
	var entry byte
	game := model.NewGame()
	for entry != EXIT { // todo finish the game, pass turn, display winner
		PrintBoard(&game)
		PrintInfo(&game)
		i, j := ReadInput(&game)
		game.Play(i, j)
	}
}

func main() {

	GameLoop()
}
