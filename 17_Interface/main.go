package main

import "fmt"



type payment struct{}


func (p payment) makePayment(amount float64) {
	// Implementation for making payment
	razorpayPaymentGw := razorpay{}
	razorpayPaymentGw.pay(amount)
}

type razorpay struct {}


func (r razorpay) pay(amount float64) {
	fmt.Println("make payment using razorpay", amount)
}

func main() {

	newPayment := payment{}
	newPayment.makePayment(1000)


}
