package main

import (
	"fmt"
	"image"
	"time"
)

func worker() {
	pos := image.Point{X: 10, Y: 10}
	direction := image.Point{X: 1, Y: 0}
	next := time.After(1500 * time.Millisecond)
	for {
		select {
		case <- next:
			pos = pos.Add(direction)
			fmt.Println("current position is", pos)
			next = time.After(1500 * time.Millisecond)
		}
	}
}

func main() {
	worker()
}
