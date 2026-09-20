package main

import "fmt"



func sum(nums ...int) int {
	total := 0

	for _, num := range nums {
		total += num
	}

	return total
}

func main() {

	result := sum(1, 2, 3, 4, 5)
	fmt.Println("The sum is:", result)

	nums := []int{10, 20, 30, 40, 50}
	result2 := sum(nums...)
	fmt.Println("The sum is:", result2)
}
