package main

import "fmt"



func changeOrderStatus(status string) {
	fmt.Println("Changing order status to", status)
}
 
func main() {
	changeOrderStatus("shipped")
}
