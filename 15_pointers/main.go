package main

import "fmt"

// by value
func changeNum(num int) {
	num = 5
	fmt.Println("In changeNum", num)
}

// by reference
func changeNumPtr(num *int) {
	*num = 5
	fmt.Println("In changeNumPtr", *num)
}

func main() {
	num := 1

	changeNum(num)
	fmt.Println("In main", num)

	changeNumPtr(&num)
	fmt.Println("In main", num)

}
