package main

import "fmt"

func main() {
	var input string
	var result, err bool
	_, _ = fmt.Scanf("%s", &input)
	switch input {
	case "true", "yes", "1":
		result = true
	case "false", "no", "0":
		result = false
	default:
		err = true
	}
	if err {
		fmt.Println("What you entered cannot be converted to a bool!")
	} else {
		fmt.Println(result)
	}
}
