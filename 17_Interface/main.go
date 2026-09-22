package main

import "fmt"

// Interface
type paymenter interface {
	pay(amount float64)
}

type payment struct{
	gateway paymenter
}


func (p payment) makePayment(amount float64) {
	// Implementation for making payment
	// razorpayPaymentGw := razorpay{}
	// razorpayPaymentGw.pay(amount)

	// stripePaymentGw := stripe{}
	// stripePaymentGw.pay(amount)
	p.gateway.pay(amount)
}

type razorpay struct {}


func (r razorpay) pay(amount float64) {
	fmt.Println("make payment using razorpay", amount)
}

type stripe struct {}

func (s stripe) pay(amount float64) {
	fmt.Println("make payment using stripe", amount)
}

func main() {

	// stripePaymentGw := stripe{}
	razorpayPaymentGw := razorpay{}

	newPayment := payment{gateway: razorpayPaymentGw}
	newPayment.makePayment(1000)


}
