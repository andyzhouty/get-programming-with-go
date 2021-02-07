package main

import "fmt"

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

func main() {
	columbiaMemorial := location{
		lat:  coordinate{14, 34, 6.2, 'S'}.decimal(),
		long: coordinate{175, 28, 21.5, 'E'}.decimal(),
	}
	challengerMemorial := location{
		lat:  coordinate{1, 56, 46.3, 'S'}.decimal(),
		long: coordinate{354, 28, 24.2, 'E'}.decimal(),
	}
	bradbury := location{
		lat:  coordinate{4, 35, 22.2, 'S'}.decimal(),
		long: coordinate{137, 26, 30.1, 'E'}.decimal(),
	}
	elysium := location{
		lat:  coordinate{4, 30, 0, 'N'}.decimal(),
		long: coordinate{135, 54, 0, 'E'}.decimal(),
	}
	fmt.Println(columbiaMemorial, challengerMemorial, bradbury, elysium)
}
