package main

import (
	"fmt"
	"slices"
)

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


	// slice equal methods
	var nums6 = []int{1, 2}
	var nums7 = []int{1, 2}
	fmt.Println(nums6 == nil)
	fmt.Println(nums7 == nil)
	fmt.Println(len(nums6) == len(nums7))
	fmt.Println(cap(nums6) == cap(nums7))
	fmt.Println(nums6[0] == nums7[0])
	fmt.Println(nums6[1] == nums7[1])
	fmt.Println(slices.Equal(nums6, nums7))


	// 2D slices
	var nums8 = [][]int{{1, 2}, {3, 4}}
	fmt.Println(nums8)
	fmt.Println(nums8[0][0])
	fmt.Println(nums8[1][1])
}
