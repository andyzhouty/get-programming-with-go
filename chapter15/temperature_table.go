package main

import "fmt"


func celsiusToFahrenheit(celsius float64) float64 {
	return float64(celsius*9.0/5.0 + 32)
}
func fahrenheitToCelsius(fahrenheit float64) float64 {
	return float64((fahrenheit - 32) * 5.0 / 9.0)
}

func drawBorder() {
	for i := 0; i < 20; i++ {
		fmt.Print("=")
	}
	fmt.Println()
}

func drawHeader(cToF bool) {
	if cToF {
		fmt.Println("| °C     | °F     |")
	} else {
		fmt.Println("| °F     | °C     |")
	}
}

func drawDataLine(temp1 float64, temp2 float64) {
	fmt.Printf("| %6.1f | %6.1f |\n", temp1, temp2)
}

func drawTable(convert func(temperature float64) float64, cToF bool) {
	drawBorder()
	drawHeader(cToF)
	drawBorder()
	for i := -40.0; i <= 100.0; i += 5 {
		converted := convert(i)
		drawDataLine(i, converted)
	}
}

func main() {
	drawTable(celsiusToFahrenheit, true)
	drawTable(fahrenheitToCelsius, false)
}
