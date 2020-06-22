package main

import "fmt"

func main() {
	/*
		Arrays
		Two types:
			a) Ones that need the size of the array and
				the type of data needs to be specified
			b) Ones that don't AKA array slices
	*/
	//Declaration
	// var fruitArr [2]string

	//Assignment
	// fruitArr[0] = "Mango"
	// fruitArr[1] = "Banana"

	//Shorthand for declaration and assignment
	fruitArr := [2]string{"Mango", "Banana"}

	//Output
	fmt.Println(fruitArr)
	fmt.Println(fruitArr[0])
	fmt.Println(fruitArr[1])

	//For array slices
	fruitSlice := []string{"Mango", "Banana", "Watermelon", "Orange"}
	fmt.Println(fruitSlice)
	//To get length of array
	fmt.Println(len(fruitSlice))
	//To output a specific range of values
	fmt.Println(fruitSlice[1:3])

}
