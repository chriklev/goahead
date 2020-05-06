package main

import (
	"fmt"
	"gobot/gogame"
	"strconv"
)

func printBoard(board [][]int8) {
	letters := "      A B C D E F G H I J K L M N O P Q R S"
	fmt.Println()
	fmt.Println(letters)

	for i := range board {
		fmt.Printf("%5.d", i+1)

		for j := range board[i] {
			point := board[j][i]
			if point == 0 {
				if j != 0 {
					fmt.Print("─┼")
				} else {
					fmt.Print(" ┼")
				}
			} else if point == 1 {
				fmt.Print("⚫")
			} else if point == -1 {
				fmt.Print("⚪")
			}
		}
		fmt.Printf(" %.d\n", i+1)
	}
	fmt.Println(letters)
	fmt.Println()
}

func main() {

	game := gogame.Game{}
	game.NewGame(19)

	game.MakeMove(4, 3)
	game.MakeMove(4, 4)
	game.MakeMove(4, 5)
	game.MakeMove(7, 5)
	game.MakeMove(3, 4)
	game.MakeMove(8, 5)

	for {
		printBoard(game.Board)
		if game.TurnColor == 1 {
			fmt.Print(" ⚫ -> ")
		} else {
			fmt.Print(" ⚪ -> ")
		}
		var input string
		fmt.Scan(&input)
		x := input[0] - 97
		y, err := strconv.Atoi(input[1:])
		if err != nil {
			panic(err)
		}
		y--

		game.MakeMove(int(x), y)
	}
}
