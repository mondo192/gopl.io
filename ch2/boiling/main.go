package main

import "fmt"

const farenheitBoilingPoint = 212.0

func boiling() {
	var farenheit = farenheitBoilingPoint
	var celsius = (farenheit - 32) * 5 / 9
	fmt.Printf("Boiling point = %g fahrenheit or %g celsius\n", farenheit, celsius)
}

func main() {
	boiling()
}
