package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var piggyBank float64
	for piggyBank <= 20.0 {
		depositId := rand.Intn(3)
		var deposit float64
		switch depositId {
		case 0:
			deposit = 0.05
		case 1:
			deposit = 0.10
		case 2:
			deposit = 0.25
		}
		piggyBank += deposit
		fmt.Printf("Deposit %.2f, Total %.2f\n", deposit, piggyBank)
	}
}
