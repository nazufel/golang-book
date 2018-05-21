package main

import "fmt"

func main() {
	fmt.Print("Enter a length in feet to convert to meters: ")
	var input float64
	fmt.Scanf("%f", &input)

	meters := input * 0.3048

	fmt.Println(input, "in feet is", meters, "in meters.")
}
