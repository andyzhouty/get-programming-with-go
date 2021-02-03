package main

import "fmt"

// terraform函数没有实际效果
func terraform(planets [8]string) {
	for i := range planets {
		planets[i] = "New " + planets[i]
	}
}
func main() {
	planets := [...]string{
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupiter",
		"Saturn",
		"Uranus",
		"Neptune",
	}
	terraform(planets)
	fmt.Println(planets)
}
