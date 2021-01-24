package main

import "fmt"

func main() {
    for _, c := range "L fdph, L vdz, L frqtxhuhg." {
    	if c >= 'a' && c <= 'z' {
			c -= 3
		}
    	fmt.Printf("%c", c)
	}
}
