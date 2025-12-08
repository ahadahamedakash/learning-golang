package main

import "fmt"

type paymenter interface {
	// pay(amount float32) bool
	pay(amount float32)
}

type payment struct {
	gateway paymenter
}

// METHOD IN PAYMENT
// open close principle (open for extension & close for modification)
func (p payment) makePaymen(amount float32){
	// INSTANCE
	// razorpayPaymentGW := razorpay{}
	// razorpayPaymentGW.pay(amount)

	p.gateway.pay(amount)
	
}

type razorpay struct {}

// METHOD IN RAZORPAY
func (r razorpay) pay(amount float32) {
	// LOGIC TO MAKE PAYMENT
	fmt.Println("MAKING PAYMENT USING RAZORPAY: ", amount)
}

type stripe struct {}

func (s stripe) pay(amount float32){
	fmt.Println("MAKING PAYMENT USING STRIPE: ", amount)
}

func main(){

	stripePaymentGW := stripe{}

	newPayment := payment {
		gateway: stripePaymentGW,
	}

	newPayment.makePaymen(100)
}
