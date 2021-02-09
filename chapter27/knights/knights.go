package main

import "fmt"

type item struct {
	name string
}

type character struct {
	name string
	leftHand item
	itemValid bool
}

func (c *character) pickup(i *item) {
	c.leftHand = *i
	c.itemValid = true
	fmt.Printf("%v picked up %v.\n", c.name, i.name)
}

func (c *character) give(to *character) {
	if to == nil || !c.itemValid{
		return
	}
	to.leftHand = c.leftHand
	fmt.Printf("%v gave %v to %v.\n", c.name, to.leftHand.name, to.name)
}

func main() {
	arthur := character{name: "Arthur"}
	knight := character{name: "Knight"}

	shovel := item{name: "a shovel"}
	arthur.pickup(&shovel)
	arthur.give(&knight)
}
