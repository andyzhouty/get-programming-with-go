package main

import (
	"fmt"
	"math"
)

type world struct {
	radius float64
}

type location struct {
	lat, long float64
}

type coordinate struct {
	d, m, s float64
	h       rune
}

func (c coordinate) decimal() float64 {
	sign := 1.0
	switch c.h {
	case 'S', 'W', 's', 'w':
		sign = -1
	}
	return sign * (c.d + c.m/60 + c.s/3600)
}

func rad(deg float64) float64 {
	return deg * math.Pi / 180
}

func (w world) distance(p1, p2 location) float64 {
	s1, c1 := math.Sincos(rad(p1.lat))
	s2, c2 := math.Sincos(rad(p2.lat))
	clong := math.Cos(rad(p1.long - p2.long))
	return w.radius * math.Acos(s1*s2+c1*c2*clong)
}

var earth = world{radius: 6371.0}
var mars = world{radius: 3389.5}

func main() {
	london := location{
		lat:  coordinate{51, 30, 0, 'N'}.decimal(),
		long: coordinate{0, 8, 0, 'W'}.decimal(),
	}
	paris := location{
		lat:  coordinate{48, 51, 0, 'N'}.decimal(),
		long: coordinate{2, 21, 0, 'E'}.decimal(),
	}
	fmt.Printf("Distance between London and Paris is %vkm.\n", earth.distance(london, paris))
	shanghai := location{
		lat:  coordinate{31, 16, 58, 'N'}.decimal(),
		long: coordinate{121, 51, 31, 'E'}.decimal(),
	}
	beijing := location{
		lat:  coordinate{39, 54, 26.37, 'N'}.decimal(),
		long: coordinate{116, 23, 29.22, 'E'}.decimal(),
	}
	fmt.Printf("Distance between Shanghai and Beijing is %vkm.\n", earth.distance(shanghai, beijing))
	mountSharp := location{
		lat:  coordinate{5, 4, 48, 'S'}.decimal(),
		long: coordinate{137, 51, 0, 'E'}.decimal(),
	}
	mountOlympus := location{
		lat:  coordinate{18, 39, 0, 'N'}.decimal(),
		long: coordinate{226, 12, 0, 'E'}.decimal(),
	}
	fmt.Printf("Distance between Mount Sharp and Mount Olympus on Mars is %vkm.\n", mars.distance(mountSharp, mountOlympus))
}
