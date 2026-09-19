package main

import "fmt"

// for -> only construct in Go for loops
func main() {

	// while loop
	// i := 1
	// for i<= 10 {
	// 	fmt.Println(i)
	// 	i = i + 1
	// }

	// infinite loop
	// for {
	// 	fmt.Println("1")
	// }


	// for loop

	// for i := 1; i <= 10; i++ {

	// 	if i == 5 {
	// 		// break
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }

	// range loop

	for i := range 5 {
		fmt.Println(i)
	}



}
