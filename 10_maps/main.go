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


	// length of a map
	fmt.Println("Length of map:", len(m))

	// deleting a key-value pair from a map
	delete(m, "age")
	fmt.Println("After deleting age:", m)


	// clear a map

	_, ok := m["age"]
	if !ok {
		fmt.Println("Key 'age' does not exist in the map.")
	} else {
		fmt.Println("Key 'age' exists in the map.")
	}

}
