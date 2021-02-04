package main

import "fmt"

type Planets []string

func (planets Planets) terraform() {
	for i := range planets {
		planets[i] = "New " + planets[i]
	}
}

func main() {
	planets := Planets([]string{
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupiter",
		"Saturn",
		"Uranus",
		"Neptune",
	})
	planets.terraform()
	fmt.Println(planets)
}
