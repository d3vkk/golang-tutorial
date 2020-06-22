package main

//How to import multiple packages and your own packages
import (
	"fmt"
	"math"

	"bitbucket.org/d3vkk/go_crash_course/03_packages/strutil"
)

func main() {
	//To round it down
	fmt.Println(math.Floor(2.7))
	//To round it up
	fmt.Println(math.Ceil(2.7))
	//Square root
	fmt.Println(math.Sqrt(16))
	fmt.Println(strutil.Reverse("racecar"))
}
