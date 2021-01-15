package main

import "fmt"

func main() {
	var (
		distance = 56000000
		speed = 100800
	) // 等价于 var distance, speed = 56000000, 100800
	fmt.Println(distance, speed)
	var weight = 149.0
	weight = weight * 0.3783 // 等价于 weight *= 0.3783
	fmt.Println(weight)
	var age = 41
	age = age + 1 // 等价于 age += 1 或 age++
	fmt.Println(age)
}
