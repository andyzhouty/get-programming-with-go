package main

import "fmt"

type turtle struct {
	x, y int
}

func (t *turtle) String() string {
	return fmt.Sprintf("(%d, %d)", t.x, t.y)
}

func (t *turtle) goUp() {
	t.y++
}

func (t *turtle) goDown() {
	t.y--
}

func (t *turtle) goRight() {
	t.x++
}

func (t *turtle) goLeft() {
	t.x--
}

func main() {
	turtle := &turtle{0, 0}
	fmt.Printf("Init position: %v\n", turtle)

	turtle.goUp()
	fmt.Printf("Go up: %v\n", turtle)

	turtle.goDown()
	fmt.Printf("Go down: %v\n", turtle)

	turtle.goLeft()
	fmt.Printf("Go left: %v\n", turtle)

	turtle.goRight()
	fmt.Printf("Go right: %v\n", turtle)
}
