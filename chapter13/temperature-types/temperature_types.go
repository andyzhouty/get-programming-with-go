package main

import "fmt"

type celsius float64
type kelvin float64

// kelvinToCelsius converts °K to °C
func kelvinToCelsius(k kelvin) celsius {
	return celsius(k - 273.15) // 此处必须进行类型转换
}
func main() {
	var k kelvin = 294.0
	c := kelvinToCelsius(k)
	fmt.Print(k, "°K is ", c, "°C")
}
