package main

import "fmt"

func main() {
	var tempK float64
	var tempC float64

	fmt.Print("Digite a temperatura em Kelvin: ")
	fmt.Scanln(&tempK)

	tempC = tempK - 273.15

	fmt.Printf("%.2f Kelvin equivalem a %.2f graus Celsius\n", tempK, tempC)
}
