package main

import "fmt"

func fahrenheitToCelsius(fahrenheit float32) float32 {
	return (fahrenheit - 32) * 5 / 9
}

func main() {
	const freezingInFahrenheit, farenheitBoilingPoint = 32.0, 212.0
	fmt.Printf("%.2fF = %.2fC\n", freezingInFahrenheit, fahrenheitToCelsius(freezingInFahrenheit))
	fmt.Printf("%.2fF = %.2fC\n", farenheitBoilingPoint, fahrenheitToCelsius(farenheitBoilingPoint))
}