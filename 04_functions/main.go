package main

import "fmt"

/*
	func greeting(name string) string{

	"greeting" - name of function
	"name" - argument
	"string" - type of the argument
	") string {" - return type (can be absent if there's no return type)

*/

func greeting(name string) string {
	//For concatenation use '+' like Javascript
	return "Hello " + name
}
func getSum(num1 int, num2 int) int {
	// For a shorter format ,
	//		func getSum(num1, num2 int) int {}
	return num1 + num2
}
func main() {
	fmt.Println(greeting("Mark"))
	fmt.Println(getSum(1, 2))
}
