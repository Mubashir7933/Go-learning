package main

import "fmt"

// interfaces are a collection of method signatures
type paymenter interface {
	cashpayment(amount float32)
}

type payment struct {
	gateway paymenter
}

func (p payment) paymentMethod(amount float32) {
	//println("Payment of amount:", amount)
	p.gateway.cashpayment(amount)
}

type paypal struct{}

func (p paypal) cashpayment(amount float32) {
	fmt.Println("Paypal payment of amount:", amount)
}

type razorpay struct{}

func (r razorpay) cashpayment(amount float32) {
	fmt.Println("Razorpay payment of amount:", amount)
}

func main() {
	paypayGw := paypal{}
	//razorGw := razorpay{}
	payingAmount := payment{
		gateway: paypayGw,
	}
	payingAmount.paymentMethod(500.00)

}
