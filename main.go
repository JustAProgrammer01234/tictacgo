package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
		board[col][row] = val
		return true
	}
	return false
}

func main() {
	board := [][]string{{" ", " ", " "}, {" ", " ", " "}, {" ", " ", " "}}
	for {
		PrintBoard(board)

		if BoardFull(board) {
			fmt.Println("No one won :(")
			break
		} else {
			winner := BoardStatus(board)
			if winner != "" {
				if winner == "X" {
					fmt.Println("X wins!")
				} else {
					fmt.Println("O wins!")
				}
				break
			}
		}

		col_scanner := bufio.NewScanner((os.Stdin))
		fmt.Printf("Enter column: ")
		col_scanner.Scan()
		col, col_err := strconv.Atoi(col_scanner.Text())

		if col_err != nil {
			break
		}

		row_scanner := bufio.NewScanner((os.Stdin))
		fmt.Printf("Enter row: ")
		row_scanner.Scan()
		row, row_err := strconv.Atoi(row_scanner.Text())

		if row_err != nil {
			break
		}

		val_scanner := bufio.NewScanner((os.Stdin))
		fmt.Printf("Enter value: ")
		val_scanner.Scan()
		val := val_scanner.Text()

		if val == "X" || val == "O" {
			if !(InsertToBoard(val, col, row, board)) {
				fmt.Println("That position is taken.")
			}
		} else {
			fmt.Println("Value must be X or O.")
		}
	}
}
