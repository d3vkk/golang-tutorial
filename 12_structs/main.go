package main

import (
	"fmt"
	"strconv"
)

/*
	OOP in Golang
	a struct is a class
	two types of methods:
		- value receivers (when NOT changing values of an object)
		- pointer receivers (when changing values of an object)
*/

// Define class
type Person struct {
	// firstName string
	// lastName  string
	// city      string
	// gender    string
	// age       int

	//Shorthand declaration
	firstName, lastName, city, gender string
	age                               int
}

// Define value receiver method
/*
	func (p Person) greet{} string  {

	"p" - is changeable and represents the 'this' keyword
	in OOP languages like Javascript & PHP
	"Person" - name of class
	"greet()" - name of value receiver method
	"string {" - return type for the method
*/

func (p Person) greet() string {
	//age has to be converted from integer to string
	return "Hello, am " + p.firstName + " " + p.lastName + ". I am " + strconv.Itoa(p.age) + " years old."
}

// Define pointer receiver method
// asterisk is used before class name
func (p *Person) increaseAge() string {
	p.age++
	return "Hello, am " + p.firstName + " " + p.lastName + ". I am " + strconv.Itoa(p.age) + " years old."
}

// Define another pointer receiver method
func (p *Person) getMarried(spouseLastName string) string {
	if p.gender == "f" {
		p.lastName = spouseLastName
	}
	return p.lastName
}

func main() {
	//Define object
	// person1 := Person{firstName: "Mark", lastName: "Matano", city: "Windhoek", gender: "m", age: 30}
	//Shorthand assignment but less readable
	person1 := Person{"Mark", "Matano", "Windhoek", "m", 30}

	fmt.Println(person1)
	fmt.Println(person1.firstName)

	//Changing value
	person1.age++
	fmt.Println((person1.age))

	//Calling a method
	fmt.Println(person1.greet())
	fmt.Println(person1.increaseAge())

	person2 := Person{"Sandy", "Nutwa", "Kinsasha", "f", 29}
	person2.getMarried(person1.lastName)
	fmt.Println(person2)
}
