package main

import "fmt"

func main() {
	/*
		A map is a key value pair
		like JSON in Javascript or Libraries in Python
	*/
	//Declaration
	// nameofmap := make(map[typeforthekey]typeforthevalue)
	// emails := make(map[string]string)

	//Assignment
	// emails["Mark"] = "mark@hotmail.com"
	// emails["Julia"] = "julia@hotmail.com"
	// emails["Eric"] = "eric@hotmail.com"

	//Shorthand for declaration and assignment
	emails := map[string]string{"Mark": "mark@hotmail.com", "Julia": "julia@hotmail.com"}
	//You can add another map later
	emails["Eric"] = "eric@hotmail.com"

	//Output
	fmt.Println(emails)
	fmt.Println(emails["Mark"])

	//To get length
	fmt.Println(len(emails))

	//Delete
	delete(emails, "Mark")
	fmt.Println(emails)
}
