package main

import "fmt"


func add(a int, b int) int {
	return a + b
}

// 2nd method 
func add2(a, b int) int {
	return a + b
}

func getLanguages() (string, string, string) {
	return "golang", "python", "java"
}



func processIt(fn func(a int) int) {
	fn(10)
}

func main() {
	result := add(5, 3)
	fmt.Println(result)

	result2 := add2(10, 20)
	fmt.Println(result2)

	language1, language2, language3 := getLanguages()
	fmt.Println(language1, language2, language3)
	fmt.Println(getLanguages())


	fn := func(a int) int {
		return a * 2
	}

	processIt(fn)
}

