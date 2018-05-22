package main

import (
	"fmt"
	"sort"
)

func findLarge(nums ...int) int {
	// sort the list
	sort.Ints(nums)

	// assign var highest int, the last value in the sorted array.
	// [len(nums)-1] subtracts one because lists start at 0
	highest := nums[len(nums)-1]

	// print the hightest number in the list
	fmt.Println(highest)

	// go-plus wants me to return at the end. not sure why, but it compiles.
	return highest
}

func main() {
	// define list of numbers
	nums := []int{48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17}
	// call the findLarge func and feed it an endless list of inputs
	findLarge(nums...)
}
