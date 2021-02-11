package main

import "fmt"

var valuesSet []string

func filterInputs(downstream chan string, values []string) {
	for _, v := range values {
		var duplicated bool
		for _, value := range valuesSet {
			if value == v {
				duplicated = true
			}
		}
		if !duplicated {
			valuesSet = append(valuesSet, v)
			downstream <- v
		}
	}
	close(downstream)
}

func printSet(upstream chan string) {
	for v := range upstream {
		fmt.Println(v)
	}
}

func main() {
	c := make(chan string)
	values := []string{"one", "two", "one", "three"}
	go filterInputs(c, values)
	printSet(c)
}
