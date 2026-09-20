package main

import "fmt"

func main() {

	// creating a map

	m := make(map[string]string)

	// setting values in a map
	m["name"] = "John"
	m["age"] = "30"
	m["city"] = "New York"

	fmt.Println(m)
	fmt.Println(m["name"])
	

	// getting values from a map
	name := m["name"]
	age := m["age"]
	city := m["city"]

	println(name)
	println(age)
	println(city)
}
