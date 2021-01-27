package main

import "C"
import "fmt"

func kelvinToCelsius(k float64) float64 {
	return k - 273.15
}
func celsiusToFahrenheit(c float64) float64 {
	return c*9.0/5.0 + 32.0
}
func kelvinToFahrenheit(k float64) float64 {
	c := kelvinToCelsius(k)
	f := celsiusToFahrenheit(c)
	return f
}
func main() {
	kelvin := 233.0
	fmt.Printf("%.0f°K is %.2f°C.\n", kelvin, kelvinToCelsius(kelvin))
	kelvin = 0
	fmt.Printf("%.0f°K is %.2f°F.\n", kelvin, kelvinToFahrenheit(kelvin))
}
