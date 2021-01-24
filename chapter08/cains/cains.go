package main

import "fmt"

func main() {
	const distance = 236000000000000000
	const lightSpeed = 299792
	const kmToLy = lightSpeed * 86400 * 365
	const distanceLy = distance / kmToLy
	fmt.Println("Cains Major Dwarf is", distanceLy, "light years away.")
}
