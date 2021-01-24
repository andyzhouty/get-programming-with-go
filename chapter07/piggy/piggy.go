package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var piggyBank int
	for piggyBank <= 2000 {
		depositId := rand.Intn(3)
		var deposit int
		switch depositId {
		case 0:
			deposit = 5
		case 1:
			deposit = 10
		case 2:
			deposit = 25
		}
		piggyBank += deposit
		dollars := piggyBank / 100
		cents := piggyBank % 100
		fmt.Printf("Current total: $%d.%.2d\n", dollars, cents)
	}
}
