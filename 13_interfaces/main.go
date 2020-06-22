package main

import (
	"fmt"
	"math"
)

/*
	OOP in Golang
	an interface is a class that only declares empty functions that must all be defined when being implemented by a class
*/

//Define interface
type Shape interface {
	//declare empty function
	// that will get the area of different shapes
	area() float64
}

type Circle struct {
	x, y, radius float64
}
type Rectangle struct {
	width, height float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (r Rectangle) area() float64 {
	return r.width * r.height
}

// You can also declare a method for an interface
// to take in the shape
func getArea(s Shape) float64 {
	return s.area()
}

func main() {
	circle1 := Circle{3, 5, 6}
	fmt.Println(circle1.area())
	fmt.Println(getArea(circle1))

	rect1 := Rectangle{3, 5}
	fmt.Println(rect1.area())
	fmt.Println(getArea(rect1))

	/*
		After refactoring,
		you can remove the variable declarations,
		as they are only used once
	*/
	fmt.Println(Circle{3, 5, 6}.area())
	fmt.Println(Rectangle{3, 5}.area())
}
