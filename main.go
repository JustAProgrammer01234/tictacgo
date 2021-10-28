package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func PrintBoard(board [][]string) {
	for _, col := range board {
		fmt.Println(strings.Join(col, " | "))
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

func GetPos(col bool) int {
	var messageInput string
	var err string
	if col {
		messageInput = "Enter column: "
		err = "Column must be a number."
	} else {
		messageInput = "Enter row: "
		err = "Row must be a number."
	}
	for {
		pos_scanner := bufio.NewScanner((os.Stdin))
		fmt.Printf(messageInput)
		pos_scanner.Scan()
		c, col_err := strconv.Atoi(pos_scanner.Text())

		if col_err != nil {
			fmt.Println(err)
			continue
		} else {
			return c
		}
	}
}

func GetValue() string {
	for {
		val_scanner := bufio.NewScanner((os.Stdin))
		fmt.Printf("Enter value: ")
		val_scanner.Scan()
		val := val_scanner.Text()
		if val == "X" || val == "O" {
			return val
		} else {
			fmt.Println("Value must be X or O.")
		}
	}
}

func main() {
	fmt.Println("Tictacgo")
	board := [][]string{{" ", " ", " "}, {" ", " ", " "}, {" ", " ", " "}}
	for {
		fmt.Println("")
		PrintBoard(board)
		fmt.Println("")

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

		col := GetPos(true)
		row := GetPos(false)
		val := GetValue()

		if !(InsertToBoard(val, col, row, board)) {
			fmt.Println("That position is taken.")
		}
	}
}
