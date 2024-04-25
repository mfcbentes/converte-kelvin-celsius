package main

import "fmt"

func main() {
	tempK := 120.00
	var tempC float64

	tempC = tempK - 273.15

	fmt.Printf("%.2f Kelvin equivalem a %.2f graus Celsius \n", tempK, tempC)
}
