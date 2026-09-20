package main

import "fmt"


func add(a int, b int) int {
	return a + b
}

// 2nd method 
func add2(a, b int) int {
	return a + b
}

func main() {
	result := add(5, 3)
	fmt.Println(result)

	result2 := add2(10, 20)
	fmt.Println(result2)
}

