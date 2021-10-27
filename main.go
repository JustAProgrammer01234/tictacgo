package main

import (
	"fmt"
	"strings"
)

func PrintBoard(board [][]string) {
	for _, row := range board {
		fmt.Println(strings.Join(row, " | "))
	}
}

func BoardStatus(board [][]string) string {
	for col := 0; col < len(board); col++ {
		if board[col][0] == "X" && board[col][1] == "X" && board[col][2] == "X" {
			return "X"
		} else if board[col][0] == "O" && board[col][1] == "O" && board[col][2] == "O" {
			return "O"
		}
	}
	for row := 0; row < len(board); row++ {
		if board[0][row] == "X" && board[1][row] == "X" && board[2][row] == "X" {
			return "X"
		} else if board[0][row] == "O" && board[1][row] == "O" && board[2][row] == "O" {
			return "O"
		}
	}
	if (board[0][0] == "X" && board[1][1] == "X" && board[2][2] == "X") ||
		(board[2][0] == "X" && board[1][1] == "X" && board[0][2] == "X") {
		return "X"
	} else if (board[0][0] == "O" && board[1][1] == "O" && board[2][2] == "O") ||
		(board[2][0] == "O" && board[1][1] == "O" && board[0][2] == "O") {
		return "O"
	}
	return ""
}

func BoardFull(board [][]string) bool {
	fullRows := 0
	for _, col := range board {
		if col[0] != " " && col[1] != " " && col[2] != " " {
			fullRows++
		}
	}
	if fullRows < 3 {
		return false
	}
	return true
}

func InsertToBoard(val string, col int, row int, board [][]string) bool {
	if board[col][row] == " " {
		return false
	}
	board[col][row] = val
	return true
}

func main() {
	board := [][]string{{" ", " ", " "}, {" ", " ", " "}, {" ", " ", " "}}
	PrintBoard(board)
}
