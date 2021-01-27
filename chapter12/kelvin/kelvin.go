package main

import "fmt"

// kelvinToCelsius会把开尔文转换为摄氏度
func kelvinToCelsius(k float64) float64 {
	k -= 273.15
	return k
}
func main() {
	kelvin := 294.0
	celsius := kelvinToCelsius(kelvin)
	fmt.Print(kelvin, "°K is ", celsius, "°C")
}
