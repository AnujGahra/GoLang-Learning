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

	var nums2 = make([]int, 0, 5)
	fmt.Println(nums2)
	// capacity of slice is 5
	fmt.Println(cap(nums2))

	// slice initialization
	nums3 := []int{1, 2, 3}
	fmt.Println(nums3)

	nums2 = append(nums2, 1)
	nums2 = append(nums2, 2)
	fmt.Println(nums2)
	fmt.Println(cap(nums2))


	// copying slices
	nums4 := make([]int, len(nums3))
	copy(nums4, nums3)
	fmt.Println(nums4)

	// slicing slices
	nums5 := nums3[1:3]
	fmt.Println(nums5)


	
}
