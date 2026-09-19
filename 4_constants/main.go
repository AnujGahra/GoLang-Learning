package main


import "fmt"

// const name string = "John Doe"

func main() {
	const name string = "John Doe"
	fmt.Println(name)

	
	const (
		port = 8080
		host = "localhost"
	)
	
	fmt.Println("Server is running on", host, ":", port)
}




