package main

import "fmt"

// slices are like arrays but they are dynamic in size. They are built on top of arrays and provide more functionality. Slices are more flexible than arrays and can be resized, appended to, and copied.
// + useful methods: append, copy, len, cap

func main() {

	// uninitialized slice is nil
	var nums []int
	fmt.Println(nums)

	fmt.Println(len(nums))
	fmt.Println(nums == nil)

	var nums2 = make([]int, 2)
	fmt.Println(nums2)
}
