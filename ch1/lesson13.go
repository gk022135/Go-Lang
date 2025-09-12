package ch1

import "fmt"

//learning the interfaces in GO

type payment struct {}


func (p payment) makePayment(amount float32){
	razorPayment := razorpay{}
	razorPayment.pay(amount)

}


type razorpay struct {}

func (r razorpay) pay(amount float32){
	// logic to make payment

	fmt.Println("making payment suing razor pay")
}


func main(){
	newPayment := payment{}
	newPayment.makePayment(100)
}