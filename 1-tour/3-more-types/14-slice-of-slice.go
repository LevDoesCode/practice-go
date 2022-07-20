package main

import (
	"fmt"
	"strings"
)

func main() {
	// Create a tic-tac-toe board.
	// The slice can be of any type
	board := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}

	test := [][]int{{1, 2, 3}, {2, 3, 4}} // int slice
	fmt.Println(test)
	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}
