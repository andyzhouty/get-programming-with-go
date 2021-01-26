package main

import (
	"fmt"
	"strings"
)

func main() {
	plainText := "wedigyouluvthegophers"
	keyword := "GOLANG"
	plainText = strings.ToUpper(plainText)
	for i, c := range plainText {
		var bridge int
		bridge = int(c - 65)
		bridge += int(keyword[i % 6] - 65)
		bridge += 65
		if bridge > 90 {
			bridge -= 26
		}
		fmt.Printf("%c", bridge)
	}
}
