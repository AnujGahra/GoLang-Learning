package main

import "time"


type order struct {
	id string
	amount float64
	status string
	createdAt time.Time // nanoseconds precision
}

func main() {

	order1 := order{
		id: "1",
		amount: 100.0,
		status: "pending",
		createdAt: time.Now(),
	}

}
