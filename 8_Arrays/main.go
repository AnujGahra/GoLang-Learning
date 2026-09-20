package main

import "fmt"



func main() {

	var nums [4]int

	nums[0] = 10
	fmt.Println(nums[0])

	fmt.Println(nums)

	// array lenghth
	fmt.Println(len(nums))


	// boolean array
	var vals [4]bool
	vals[2] = true
	fmt.Println(vals)

	// string array
	var names [4]string
	names[0] = "John"
	names[1] = "Doe"
	fmt.Println(names)

	// array initialization
	var arr = [3]int{1, 2, 3}
	fmt.Println(arr)

	nums2 := [4]int{1, 2, 3, 4}
	fmt.Println(nums2)

	// array with ellipsis
	arr2 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(arr2)


	// 2D array
	var matrix [2][3]int
	matrix[0][0] = 1
	matrix[0][1] = 2
	matrix[0][2] = 3
	matrix[1][0] = 4
	matrix[1][1] = 5
	matrix[1][2] = 6
	
	fmt.Println(matrix)

	var matrix2 = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(matrix2)

	matrix3 := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(matrix3)
}
