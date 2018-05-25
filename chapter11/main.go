package main

import (
	"fmt"

	"github.com/ryanalexanderross/math"
)

func main() {
	xs := []float64{1, 2, 3, 4, 5, 0}
	avg := math.Average(xs)
	fmt.Println(avg)
	max := math.Max(xs)
	fmt.Println(max)
	min := math.Min(xs)
	fmt.Println(min)
}
