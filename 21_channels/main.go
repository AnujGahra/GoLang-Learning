package main

import (
	"fmt"
	"math/rand"
	"time"
)


func processNum(numChan chan int) {
	
	for num := range numChan {
		
		fmt.Println("Processing number", num)
		time.Sleep(time.Second)

	}

}



func main() {

	numChain := make(chan int)

	go processNum(numChain)

	// numChain <- 5

	for {
		numChain <- rand.Intn(100)
	}






	// messageChan := make(chan string) // create channel


	// messageChan <- "ping" // channel is blocking


	// msg := <-messageChan

	// fmt.Println(msg)




}