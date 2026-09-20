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

func ( o *order) getAmount() float64 {
	return o.amount
}

func ( o *order) getStatus() string {
	return o.status
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


	order2 := order{
		id: "2",
		amount: 200.0,
		status: "completed",
		createdAt: time.Now(),
	}

	fmt.Println("Order Struct", order2)
	fmt.Println("Order ID:", order2.id)
	fmt.Println("Order Amount:", order2.amount)
	fmt.Println("Order Status:", order2.status)
	fmt.Println("Order Created At:", order2.createdAt)

	fmt.Println("Order Amount:", myOrder.getAmount())
	fmt.Println("Order Status:", myOrder.getStatus())

}
