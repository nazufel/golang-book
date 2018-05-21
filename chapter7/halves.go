package main

import "fmt"

func main() {
	var x int = 20

	var half int = x / 2
	if half%2 == 0 {
		fmt.Println(half, true)
	}
	if half%2 != 0 {
		fmt.Println(half, false)
	}
}
