package main

import "fmt"

func main() {
	/*
		A pointer (i.e. '&') is used to indicate the memory address of a variable's value

		Use cases for them would be passing in alot of data values,
		Instead of passing values, passing the memory addresses of
		those values would increase performance
	*/
	a := 5
	b := &a
	fmt.Println(a, b)

	//For the types of each
	fmt.Printf("%T\n", a) //int
	fmt.Printf("%T\n", b) //*int - the asterisk shows it's a pointer

	//Use asterisk to read variable's value from memory address
	fmt.Println(*b)
	fmt.Println(*&a)

	//Change variable's value with pointer
	// since b = &a, *b = *&a
	// which is the same thing as declaring a = 10
	*b = 10
	fmt.Println(a)

}
