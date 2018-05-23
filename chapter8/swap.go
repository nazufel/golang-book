package main

import "fmt"

func swap(x, y *int) {
	*x = 2
	*y = 1
}

func main() {
	x := 1
	y := 2

	fmt.Println("Unswapped: ", x, y)

	swap(&x, &y)

	fmt.Println("Swapped: ", x, y)
}
