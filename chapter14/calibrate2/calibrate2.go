package main

import (
	"fmt"
	"math/rand"
)

type kelvin float64
type sensor func() kelvin

func fakeSensor() kelvin {
	return kelvin(rand.Intn(151) + 150)
}

func realSensor() kelvin {
	return 0
}

func calibrate(s sensor, offset kelvin) sensor {
	return func() kelvin {
		return s() + offset
	}
}

func main() {
	var offset kelvin = 5
	sensor := calibrate(realSensor, offset)
	fmt.Println(sensor())
	offset = 10
	fmt.Println(sensor()) // should output 5 as well
	sensor = calibrate(fakeSensor, offset)
	fmt.Println(sensor())
	sensor = calibrate(fakeSensor, offset)
	fmt.Println(sensor())
}
