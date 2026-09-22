package main

import "fmt"

type OrderStatus int


const (
	Received OrderStatus = iota
	Confirmed
	prepared
	Delivered
)

func changeOrderStatus(status OrderStatus) {
	fmt.Println("Changing order status to", status)
}
 
func main() {
	changeOrderStatus(Received)
}
