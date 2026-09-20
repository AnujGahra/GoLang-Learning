package main

import (
	"fmt"
	"time"
)


type order struct {
	id string
	amount float64
	status string
	createdAt time.Time // nanoseconds precision
}

func main() {

	myOrder := order{
		id: "1",
		amount: 100.0,
		status: "pending",
		createdAt: time.Now(),
	}

	fmt.Println("Order Struct", myOrder)
	fmt.Println("Order ID:", myOrder.id)
	fmt.Println("Order Amount:", myOrder.amount)
	fmt.Println("Order Status:", myOrder.status)
	fmt.Println("Order Created At:", myOrder.createdAt)

}
