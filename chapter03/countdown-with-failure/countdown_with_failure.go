package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var count = 10
	for count > 0 {
		fmt.Println(count)
		time.Sleep(time.Second)
		count--
		if rand.Intn(100) == 0 {
			break
		}
	}
	switch count {
	case 0:
		fmt.Println("Liftoff!")
	default:
		fmt.Println("Launch Failed. :(")
	}
}
