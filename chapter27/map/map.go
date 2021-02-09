package main

import "fmt"

func main() {
	var soup map[string]int
	fmt.Println(soup == nil) // true

	measurement, ok := soup["onion"]
	if ok {
		fmt.Println(measurement)
	}

	// 下面这段代码没有任何输出
	for ingredient, measurement := range soup {
		fmt.Println(ingredient, measurement)
	}
}
