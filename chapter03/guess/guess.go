package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var number = rand.Intn(100) + 1
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Let's guess a number from 1 to 100: ")
	var guess int
	for guess != number{
		var guess = rand.Intn(100) + 1
		fmt.Printf("I guess %d. ", guess)
		if guess != number {
			fmt.Println("Incorrect.")
		} else {
			break
		}
	}
	fmt.Println("Correct!")
}
