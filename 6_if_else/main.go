package main

import "fmt"

func main() {

	age := 12

	if age >= 18 {
		fmt.Println("You are an adult.")
	} else if age >= 13 && age < 18 {
		fmt.Println("You are a teenager.")
	} else {
		fmt.Println("You are a child.")
	}



	var role = "admin"
	var hashPermission = true

	if role == " admin" && hashPermission {
		fmt.Println("You have access to the admin panel.")
	} else {
		fmt.Println("You do not have access to the admin panel.")
	}


	if age :=18; age >= 18 {
		fmt.Println("You are an adult.", age)
	} else {
		fmt.Println("You are not an adult.", age)
	}
}
