package main

/*
	MAIN TYPES OF VARIABLES
	string
	bool
	int int8 int16 int32 int64
	uint uint8 uint16 uint32 uint64 uintptr
	byte - alias for uint8
	rune - alias for int32
	float32 float64
	complex64 complex128
*/
import "fmt"

//Variable can also be declared outside as global variable
var location = "CA"

func main() {
	//Using var
	var name = "Mark"
	var age = 30
	/*
		Note that string and int variables
		can also be declared i.e.
			var name string = "Mark"
			var age int32 = 30
	*/

	//To output
	fmt.Println(name, age)

	//Using const
	const isMale = true

	/*
		Shorthand variable declaration
		can only be inside a function
	*/
	//Using float
	//Default is float64 unless specified as float32
	height := 5.11
	var feet float32 = 1.3
	//To check type of variable
	//Note that '\n' is for going to the next line
	fmt.Printf("%T", height)
	fmt.Printf("%T\n", feet)

	//To declare multiple variables in one line
	email, password := "mark@gmail.com", "mark123"
	fmt.Println(email, password)

}
