package main

import "fmt"



type payment struct{
	gateway stripe
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

	newPayment := payment{}
	newPayment.makePayment(1000)


}
