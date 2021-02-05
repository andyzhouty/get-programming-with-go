package main

import "fmt"

func lengthen(slice []int) {
	capacity := cap(slice)
	var i int
	for {
		slice = append(slice, i)
		i++
		fmt.Println(slice)
		if cap(slice) > capacity {
			fmt.Printf("capacity %d", cap(slice))
			break
		}
	}
}

func main() {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice = slice[0:4:6]
	lengthen(slice)
}
