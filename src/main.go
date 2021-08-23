package main

import (
	model "./model"
	"fmt"
)

func PrintBoard(game *model.Game) {
	var i, j uint
	for i = 0; i < model.BOARD_SIZE; i++ {
		for j = 0; j < model.BOARD_SIZE; j++ {
			fmt.Printf("| %c ", game.GetPos(i, j))
		}
		fmt.Printf("| \n")
	}
}

func main() {
	fmt.Printf("Hello, world!\n")
	game := model.NewGame()

	PrintBoard(&game)
	game.Play(1, 2)
	fmt.Printf("\n\n")
	PrintBoard(&game)

	game.Play(1, 3)
	fmt.Printf("\n\n")
	PrintBoard(&game)

	game.Play(0, 0)
	fmt.Printf("\n\n")
	PrintBoard(&game)

	game.Play(1, 1)
	fmt.Printf("\n\n")
	PrintBoard(&game)

	game.Play(0, 1)
	fmt.Printf("\n\n")
	PrintBoard(&game)

	game.Play(0, 2)
	fmt.Printf("\n\n")
	PrintBoard(&game)

	game.Play(3, 3)
	fmt.Printf("\n\n")
	PrintBoard(&game)

	game.Play(2, 2)
	fmt.Printf("\n\n")
	PrintBoard(&game)

}
