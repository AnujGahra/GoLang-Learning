package main

import "fmt"


func counter() func() int {
	var count int = 0
	return func() int {
		count++
		return count
	}
}


func main() {
	// Closure is a function value that references variables from outside its body.
	// The function may access and assign to the referenced variables; in this sense the function is "bound" to the variables.

	c := counter()
	fmt.Println(c()) // 1
	fmt.Println(c()) // 2
	fmt.Println(c()) // 3

}