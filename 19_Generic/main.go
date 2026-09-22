package main

import "fmt"



func printSlice(items []int) {
	for _, item := range items {
		fmt.Println(item)
	}
}
// Generic function to print slices of any type
func printSliceT[T int | string | bool](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}

func main() {

	// slice := []int{1, 2, 3, 4, 5}

	names := []string{"Alice", "Bob", "Charlie", "David"}
	printSliceT(names)
}

