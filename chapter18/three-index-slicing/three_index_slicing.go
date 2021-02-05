package main

import "fmt"

func main() {
	planets := []string{
		"Mercury", "Venus", "Earth", "Mars",
		"Jupiter", "Saturn", "Uranus", "Neptune",
	}
	terrestrial := planets[0:4:4] // 长度为4, 容量为4
	worlds := append(terrestrial, "Ceres")
	fmt.Println(planets) // planets此时不变

	terrestrial = planets[0:4]
	worlds = append(terrestrial, "Ceres")
	fmt.Println(planets)
	fmt.Println(worlds)
}
