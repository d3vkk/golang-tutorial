package main

import "fmt"

func main() {
	/*
		A range is used to loop through arrays and maps
	*/
	//For an array
	ids := []int{34, 56, 23, 8, 99}

	//Using the index, 'i'
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	//Not using the index, 'i'
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	//Adding id's
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum: ", sum)

	//For a map
	emails := map[string]string{"Mark": "mark@hotmail.com", "Julia": "julia@hotmail.com"}
	// k for key, v for value
	for k, v := range emails {
		fmt.Printf("%s: %s\n", k, v)
	}
	for k := range emails {
		fmt.Println("Name: " + k)
	}
}
