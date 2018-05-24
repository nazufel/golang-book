package main

import (
	"fmt"
	"math"
)

//Circle struct
type Circle struct {
	ra float64 // radius of a circle
}

//Rectangle struct
type Rectangle struct {
	h, w float64 // height and width of a rectangle
}

// Area interface
type Area interface {
	findCirArea() float64
	findRecArea() float64
}

// Parameter /Circumferance interace
type Parameter interface {
	findCirCirc() float64
	findRecPar() float64
}

// print the area of a circle
func findCirArea(c Circle) {
	fmt.Println(math.Pi * (c.ra * c.ra))
}

// print the area of a rectangle
func findRecArea(r Rectangle) {
	fmt.Println(r.h * r.w)
}

//calculate the Parameter/Circumferance
func findCirCirc(c Circle) {
	fmt.Println((math.Pi * 2) * c.ra)
}

func findRecPar(r Rectangle) {
	fmt.Println((r.h * 2) + (r.w * 2))
}

func main() {
	c1 := Circle{
		2.1,
	}

	p1 := Rectangle{
		2,
		2,
	}

	// I don't like that the print statements are in the methods, will fix later.
	findCirCirc(c1)
	findRecPar(p1)
	findCirArea(c1)
	findRecArea(p1)
}

/*
also, didn't copy from the Internet either. just watched a lot of videos on the topic,
+ and applied the concepts to this exercise.
*/
