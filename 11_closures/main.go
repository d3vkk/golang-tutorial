package main

import "fmt"

/*
	A closure function is composed of an outer function and an inner function

	Inner function has access to the outer function's
	parameters as it is within it's local scope

	This example demonstrates anonymous functions and the Fibonacci Sequence
	where the output values are recursively added with the input values.

*/

//outer function
//Declaring two functions one after another, anonymizes a function
func fibonacci() func(int) int {
	sum := 0
	//inner function
	return func(x int) int {
		sum += x
		return sum

	}

}
func main() {
	//calling the outer function
	//Has to be equated to a variable since it's anonymous
	sum := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(sum(i))
	}
}
