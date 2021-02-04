package main

import (
	"fmt"
	"strings"
)

// hyperspace函数将一处围绕行星的空间
func hyperspace(worlds []string) {
	for i := range worlds {
		worlds[i] = strings.TrimSpace(worlds[i])
	}
}
func main() {
	planets := []string{" Venus ", "Earth ", " Mars"}
	hyperspace(planets)
	fmt.Println(strings.Join(planets, ""))
}
