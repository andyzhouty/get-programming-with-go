package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type coordinate struct {
	d float64 `json:"degrees"`
	m float64 `json:"minutes"`
	s float64 `json:"seconds"`
	h rune    `json:"hemispher"`
}

func (c coordinate) decimal() float64 {
	sign := 1.0
	switch c.h {
	case 'S', 'W', 's', 'w':
		sign = -1
	}
	return sign * (c.d + c.m/60 + c.s/3600)
}

func (c coordinate) String() string {
	return fmt.Sprintf("%v°%v'%.1f\" %c", c.d, c.m, c.s, c.h)
}

func (c coordinate) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		DD  float64 `json:"decimal"`
		DMS string  `json:"dms"`
		D   float64 `json:"degrees"`
		M   float64 `json:"minutes"`
		S   float64 `json:"seconds"`
		H   string  `json:"hemisphere"`
	}{
		DD:  c.decimal(),
		DMS: c.String(),
		D:   c.d,
		M:   c.m,
		S:   c.s,
		H:   string(c.h),
	})
}

func exitOnError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	coord := coordinate{135, 54, 0, 'E'}
	bytes, err := coord.MarshalJSON()
	if err != nil {
		exitOnError(err)
	}
	fmt.Println(string(bytes))
}
