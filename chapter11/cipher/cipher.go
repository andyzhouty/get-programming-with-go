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
		var char int
		char = int(c - 65)
		char += int(keyword[i % 6] - 65)
		char += 65
		if char > 90 {
			char -= 26
		}
		fmt.Printf("%c", char)
	}
}
