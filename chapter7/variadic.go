package main

import (
	"fmt"
	"sort"
)

func findLarge(args ...int) int {
	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}

	// I know I cheated, but that's what the standard library is for.
	sort.Ints(x)
	highest := x[int(len(3))]

	return highest
}

func main() {
	fmt.Println(findLarge(48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17))
}
