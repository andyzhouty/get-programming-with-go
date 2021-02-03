package main

import (
	"fmt"
	"strings"
)

func showBoard(board [8][8]rune) {
	for _, row := range board {
		for _, item := range row {
			if !(item > 'A' && item <= 'z') {
				fmt.Print("_")
			} else {
				fmt.Print(string(item))
			}
		}
		fmt.Println()
	}
}

func main() {
	var board [8][8]rune
	var firstLine = "rnbkqbnr"
	var lastLine = strings.ToUpper(firstLine)
	for column := range firstLine {
		board[0][column] = rune(firstLine[column])
	}
	for column := range board[1] {
		board[1][column] = 'p'
	}
	for column := range board[6] {
		board[6][column] = 'P'
	}
	for column := range lastLine {
		board[7][column] = rune(lastLine[column])
	}
	showBoard(board)
}
