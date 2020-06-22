package main

import "fmt"

func main() {
	x, y := 15, 10
	/*
		Note that the conditionals don't have to
		be surrounded with brackets
		But no error will occur if they are
	*/
	//If else
	if x < y {
		//The syntax here is very similar to the C language
		// as the '%d' are placeholders for int variables
		fmt.Printf("%d is less than %d\n", x, y)
	} else {
		fmt.Printf("%d is less than %d\n", y, x)
	}

	//else if
	color := "yellow"

	if color == "yellow" {
		fmt.Printf("Color is %s but not brown", color)
	} else if color == "brown" {
		fmt.Printf("Color is %s but not yellow", color)
	} else {
		fmt.Printf("Color is %s but not yellow or brown", color)
	}

	//Switch statement
	switch color {
	case "yellow":
		fmt.Printf("Color is %s but not brown", color)
	case "brown":
		fmt.Printf("Color is %s but not yellow", color)
	default:
		fmt.Printf("Color is %s but not yellow or brown", color)

	}
}
