package main

import "fmt"

type celsius float64
type fahrenheit float64
type kelvin float64

func (c celsius) kelvin() kelvin {
	return kelvin(c + 273.15)
}
func (c celsius) fahrenheit() fahrenheit {
	return fahrenheit(c*9.0/5.0 + 32)
}
func (f fahrenheit) celsius() celsius {
	return celsius((f - 32) * 5.0 / 9.0)
}
func (f fahrenheit) kelvin() kelvin {
	return f.celsius().kelvin()
}
func (k kelvin) celsius() celsius {
	return celsius(k - 273.15)
}
func (k kelvin) fahrenheit() fahrenheit {
	return k.celsius().fahrenheit()
}

func main() {
	var f fahrenheit = -40
	fmt.Printf("%.2f°F = %.2f°K = %.2f°C\n", f, f.kelvin(), f.celsius())
	var c celsius = 100
	fmt.Printf("%.2f°C = %.2f°K = %.2f°F\n", c, c.kelvin(), c.fahrenheit())
	var k kelvin = 273.15
	fmt.Printf("%.2f°K = %.2f°C = %.2f°F\n", k, k.celsius(), k.fahrenheit())
}
