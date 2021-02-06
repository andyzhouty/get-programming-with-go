package main

import "fmt"

type location struct {
	lat, long float64
}

func main() {
	bradbury := location{-4.5895, 137.4417}
	curiosity := bradbury

	curiosity.long += 0.0106

	fmt.Println(bradbury, curiosity)
}
