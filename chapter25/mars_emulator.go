package main

import (
	"fmt"
	"math/rand"
	"time"
)

type animal interface {
	move() string
	eat() string
}

type gopher struct {
	name string
}

func (g gopher) String() string {
	return g.name
}

func (g gopher) move() string {
	return g.name + " is crawling around."
}

func (g gopher) eat() string {
	return g.name + " is eating lettuce root."
}

type giraffe struct {
	name string
}

func (g giraffe) String() string {
	return g.name
}

func (g giraffe) move() string {
	return g.name + " is stepping nearby."
}

func (g giraffe) eat() string {
	return g.name + " is eating fresh leaves."
}

type rabbit struct {
	name string
}

func (r rabbit) String() string {
	return r.name
}

func (r rabbit) move() string {
	return r.name + " is hopping."
}

func (r rabbit) eat() string {
	foods := [2]string{"spinaches", "carrots"}
	randomFood := foods[rand.Intn(2)]
	return fmt.Sprintf("%s is eating %s.", r.name, randomFood)
}

const (
	hours   = 72
	sunSet  = 18
	sunRise = 6
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	for i := 1; i <= hours; i++ {
		if i%24 > sunSet || i%24 < sunRise {
			fmt.Printf("%d: All animals are sleeping.\n", i)
			continue
		}
		fmt.Printf("%d: ", i)
		var randomAnimal animal
		switch rand.Intn(3) {
		case 0:
			randomAnimal = gopher{name: "Gopher"}
		case 1:
			randomAnimal = giraffe{name: "Giraffe"}
		case 2:
			randomAnimal = rabbit{name: "Rabbit"}
		}
		switch rand.Intn(2) {
		case 0:
			fmt.Print(randomAnimal.move())
		case 1:
			fmt.Print(randomAnimal.eat())
		}
		fmt.Println()
	}
}
