package main

import (
	"fmt"
	"math"
)

type gps struct {
	current, target location
	world
}

type location struct {
	name      string
	lat, long float64
}

type world struct {
	radius float64
}

type rover struct {
	gps
}

func (l location) description() string {
	return fmt.Sprintf("%s (Lat: %.2f, Long: %.2f)", l.name, l.lat, l.long)
}

func (w world) distance(p1, p2 location) float64 {
	s1, c1 := math.Sincos(rad(p1.lat))
	s2, c2 := math.Sincos(rad(p2.lat))
	clong := math.Cos(rad(p1.long - p2.long))
	return w.radius * math.Acos(s1*s2+c1*c2*clong)
}

func rad(deg float64) float64 {
	return deg * math.Pi / 180
}

func (g gps) distance() float64 {
	return g.world.distance(g.current, g.target)
}

func (g gps) message() string {
	return fmt.Sprintf("The distance from %s to %s is %.2fkm.", g.current.name, g.target.name, g.distance())
}

var mars = world{radius: 3389.5}

func main() {
	marsGPS := gps{
		current: location{"Bradbury Landing", -4.5895, 137.4417},
		target:  location{"Elysium Planitia", 4.5, 135.9},
		world:   mars,
	}

	curiosity := rover{gps: marsGPS}
	fmt.Println(curiosity.message())
}
