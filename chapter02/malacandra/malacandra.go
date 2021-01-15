package main

import "fmt"

func main() {
	const distance = 56_000_000
	const days = 28
	var speed = distance / (days * 24)
	fmt.Printf("%v km/h\n", speed)
}
