package main

import "fmt"


func main() {
	messageChan := make(chan string) // create channel


	messageChan <- "ping" // insert


	msg := <-messageChan

	fmt.Println(msg)




}