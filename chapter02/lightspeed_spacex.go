package main

import "fmt"

func main() {
	const spaceXSpeed = 100800 // km/h
	var distance = 96300000    // km
	fmt.Println(distance/spaceXSpeed/24, "days")
}
