package main

import "fmt"

func main() {
	fmt.Print("Enter a temperature in Farienhight to convert to Celsius: ")
	var input float64
	fmt.Scanf("%f", &input)

	celsius := (input - 32) * 5 / 9

	fmt.Println(input, " degrees Farenhight is ", celsius, " degrees Celsius.")
}
