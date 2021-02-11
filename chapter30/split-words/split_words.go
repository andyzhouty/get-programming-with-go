package main

import (
	"fmt"
	"strings"
)

func splitWords(downstream chan string, sentence string) {
	for _, w := range strings.Fields(sentence) {
		downstream <- w
	}
	close(downstream)
}

func printWords(upstream chan string) {
	for w := range upstream {
		fmt.Println(w)
	}
}

func main() {
	c := make(chan string)
	sentence := "Go is the best programming language"
	go splitWords(c, sentence)
	printWords(c)
}
