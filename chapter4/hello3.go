package main

import "fmt"

var x string = "Hello World!"

func main() {
	//	x := "Hello World!"
	fmt.Println(x)
	change()
}

func change() {
	// var x string
	const x string = "first"
	fmt.Println(x)
	//	x += ", second"
	fmt.Println(x + " second")
}
