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
}
