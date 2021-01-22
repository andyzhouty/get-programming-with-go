package main

import (
	"fmt"
	"math/rand"
)

func main() {
	const smallest = 56_000_000
	const biggest = 401_000_000
	var distance = rand.Intn(biggest - smallest + 1) + smallest
	fmt.Println(distance)
}
