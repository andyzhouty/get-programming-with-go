package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	width  = 80
	height = 15
)

type Universe [][]bool

func NewUniverse() Universe {
	u := make(Universe, height)
	for i := 0; i < height; i++ {
		u[i] = make([]bool, width)
	}
	return u
}

func (u Universe) Set(x, y int, b bool) {
	u[y][x] = b
}

func (u Universe) Get(x, y int) bool {
	return u[y][x]
}

func (u Universe) String() string {
	var b byte
	buf := make([]byte, 0, (width+1)*height)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			b = ' '
			if u[y][x] {
				b = '*'
			}
			buf = append(buf, b)
		}
		buf = append(buf, '\n')
	}
	return string(buf)
}

func (u Universe) Show() {
	fmt.Print("\x0c", u.String())
}

func (u Universe) Seed() {
	rand.Seed(time.Now().UnixNano())
	activated := width * height / 4
	for i := 0; i < activated; i++ {
		x := rand.Intn(width)
		y := rand.Intn(height)
		u[y][x] = true
	}
}

func (u Universe) Alive(x, y int) bool {
	x += width
	x %= width
	y += height
	y %= height
	return u[y][x]
}

func (u Universe) Neighbors(x, y int) int {
	var neighbors int
	for i := x-1; i <= x+1; i++ {
		for j := y-1; j <= y+1; j++ {
			if (i != x || j != y) && u.Alive(i, j) {
				neighbors++
			}
		}
	}
	return neighbors
}

func (u Universe) Next(x, y int) bool {
	neighbors := u.Neighbors(x, y)
	return neighbors == 3 || (neighbors == 2 && u.Alive(x, y))
}

func Step(a, b Universe) {
	for y := 0; y < height; y++{
		for x := 0; x < width; x++{
			b.Set(x, y, a.Next(x, y))
		}
	}
	b.Show()
}

func main() {
	a, b := NewUniverse(), NewUniverse()
	a.Seed()
	for i := 0; i < 100000; i++{
		Step(a, b)
		a, b = b, a
		time.Sleep(100 * time.Millisecond)
	}
}
