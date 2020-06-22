package main

import "fmt"

func main() {
	//Long method
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	//Short method
	for i := 1; i <= 10; i++ {
		fmt.Printf("Number %d\n", i)
	}

	//FizzBuzz Challenge
	for i := 1; i <= 100; i++ {
		//'%' - represents modulus
		if i%15 == 0 {
			fmt.Printf("FizzBuzz\n")
		} else if i%3 == 0 {
			fmt.Printf("Fizz\n")
		} else if i%5 == 0 {
			fmt.Printf("Buzz\n")
		} else {
			fmt.Println(i)
		}
	}

}
