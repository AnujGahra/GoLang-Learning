package main

import (
	"fmt"
	"time"
)

func main() {
	// simple switch statement
	// i := 5
	// switch i {
	// case 1:
	// 	fmt.Println("i is 1")
	// case 2:
	// 	fmt.Println("i is 2")
	// case 3:
	// 	fmt.Println("i is 3")
	// default:
	// 	fmt.Println("i is not 1, 2, or 3")
	// }


	// switch with multiple cases
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend!")
	default:
		fmt.Println("It's a weekday.")
	}

	// more examples of multiple switch statements



	// type switch
	whoAmI := func(i interface{}) {
		switch t := i.(type) {
		case int:
			fmt.Printf("I'm an int and my value is %d\n", i)
		case string:
			fmt.Printf("I'm a string and my value is %s\n", i)
		default:
			fmt.Printf("I don't know what type I am\n", t)
		}
	}

	whoAmI(42)
	whoAmI("hello")
	whoAmI(3.14)
}
