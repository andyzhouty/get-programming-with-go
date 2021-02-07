package main

import (
	"fmt"
	"math"
)

type world struct {
	radius float64
}

type location struct {
	name      string
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

var mars = world{radius: 3389.5}

func main() {
	columbiaMemorial := location{
		name: "Columbia Memorial Station",
		lat:  coordinate{14, 34, 6.2, 'S'}.decimal(),
		long: coordinate{175, 28, 21.5, 'E'}.decimal(),
	}
	challengerMemorial := location{
		name: "Challenger Memorial Station",
		lat:  coordinate{1, 56, 46.3, 'S'}.decimal(),
		long: coordinate{354, 28, 24.2, 'E'}.decimal(),
	}
	bradbury := location{
		name: "Bradbury Landing",
		lat:  coordinate{4, 35, 22.2, 'S'}.decimal(),
		long: coordinate{137, 26, 30.1, 'E'}.decimal(),
	}
	elysium := location{
		name: "Elysium Planitia",
		lat:  coordinate{4, 30, 0, 'N'}.decimal(),
		long: coordinate{135, 54, 0, 'E'}.decimal(),
	}
	landingSites := []location{
		columbiaMemorial, challengerMemorial, bradbury, elysium,
	}
	maxDistance := mars.distance(landingSites[0], landingSites[1])
	minDistance := mars.distance(landingSites[0], landingSites[1])
	maxIndexes := make([]int, 2, 2)
	minIndexes := make([]int, 2, 2)
	for i := range landingSites {
		for j := i + 1; j < len(landingSites); j++ {
			distance := mars.distance(landingSites[i], landingSites[j])
			if distance > maxDistance {
				maxDistance = distance
				maxIndexes = []int{i, j}
			}
			if distance < minDistance {
				minDistance = distance
				minIndexes = []int{i, j}
			}
		}
	}
	fmt.Printf(
		"Max distance is %.2fkm (between %s and %s).\n",
		maxDistance,
		landingSites[maxIndexes[0]].name,
		landingSites[maxIndexes[1]].name,
	)

	fmt.Printf(
		"Min distance is %.2fkm (between %s and %s).\n",
		minDistance,
		landingSites[minIndexes[0]].name,
		landingSites[minIndexes[1]].name,
	)
}
