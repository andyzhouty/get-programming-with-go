package main

import (
	"errors"
	"fmt"
	"os"
)

const (
	rows    = 9
	columns = 9
)

var (
	ErrRow               = errors.New("duplicated digit in the same row")
	ErrColumn            = errors.New("duplicated digit in the same column")
	ErrNearbyDigits      = errors.New("duplicated with a digit in the nearby 8")
	ErrInvalidDigit      = errors.New("each digit must be greater 0 and smaller than 9")
	ErrCannotChangeDigit = errors.New("cannot change the prebuilt digit")
)

type Grid [rows][columns]int8

func NewSudoku(digits [rows][columns]int8) Grid {
	var grid Grid
	for r := 0; r < rows; r++ {
		for c := 0; c < columns; c++ {
			if digits[r][c] < 0 || digits[r][c] > 9 {
				fmt.Println("")
				os.Exit(1)
			}
			grid[r][c] = digits[r][c]
		}
	}
	return grid
}

func (g *Grid) isSet(row, column int8) bool {
	return g[row][column] != 0
}

func (g *Grid) Set(row, column int8, digit int8) error {
	fmt.Printf("Trying to set (%v, %v) to %v\n", row, column ,digit)
	if g.isSet(row, column) {
		return ErrCannotChangeDigit
	}
	if digit < 0 || digit >= columns {
		return ErrInvalidDigit
	}
	for c := 0; c < columns; c++ {
		if g[row][c] == digit {
			return ErrRow
		}
	}
	for r := 0; r < rows; r++ {
		if g[r][column] == digit {
			return ErrColumn
		}
	}
	for r := row - 1; r <= row+1; r++ {
		for c := column - 1; c <= column+1; c++ {
			if r < 0 || c < 0 || r > 8 || c > 8 {
				continue
			}
			if g[r][c] == digit {
				return ErrNearbyDigits
			}
		}
	}
	return nil
}

func (g *Grid) Clear(row, column int8) error {
	if g.isSet(row, column) {
		return ErrCannotChangeDigit
	}
	g[row][column] = 0
	return nil
}

func main() {
	s := NewSudoku([rows][columns]int8{
		{5, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	})
	err := s.Set(0, 2, 0)
	if err != nil {
		fmt.Println(err)
	}
	err = s.Set(0, 2, 5)
	if err != nil {
		fmt.Println(err)
	}
	err = s.Set(1, 1, 3)
	if err != nil {
		fmt.Println(err)
	}
	err = s.Set(5, 1, 4)
	if err != nil {
		fmt.Println(err)
	}
	err = s.Clear(0, 0)
	if err != nil {
		fmt.Println(err)
	}
}
