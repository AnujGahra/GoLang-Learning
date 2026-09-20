package main

import "fmt"

// Iteration for data structure
func main() {


	nums := []int{1, 2, 3, 4, 5}

	for i:=0; i<len(nums); i++ {
		fmt.Println(nums[i])
	}

	for i, v := range nums {
		fmt.Println(i, v)
	}

	for _, v := range nums {
		fmt.Println(v)
	}


	m := map[string]string{
		"name": "John",
		"age": "30",
		"city": "New York",
	}

	for k, v := range m {
		fmt.Println(k, v)
	}

	for k := range m {
		fmt.Println(k)
	}

	for _, v := range m {
		fmt.Println(v)
	}	


	for i, c := range "Hello" {
		fmt.Println(i, c)
	}

	
	for i, c := range "Hello" {
		fmt.Println(i, string(c))
	}
}
