package ch1

import (
	"fmt"
	"time"
)

//learning the struct in go lang

/*
>> mutliple thing ko geoup krne ke liye
like order ke liye, student details
*/

type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time // nanoseond prescision , much accurate
}

// recever type (o order) --> ye function order struct se attached ho jayegi
func (o *order) changesStatus(status string) {
	o.status = status
}

func (o *order) changesAmount(amount float32) {
	o.amount = amount
}

//construtor with go

func newOrder(id string, amount float32, status string) *order {

	//inital setups goes here .....
	order12 := order{
		id:     id,
		amount: amount,
		status: status,
	}
//constructor my address return krte hai
	return &order12
}

func main() {
	myorder := newOrder("1", 23.34, "paid")
	fmt.Println("myorder ",myorder)
	
	// var order order =

	//now create the instances of the struct

	order1 := order{
		id:     "1",
		amount: 50.23,
		status: "paid",
	}

	order2 := order{
		id:     "1ad",
		amount: 504.23,
		status: "unpaid",
	}

	order1.createdAt = time.Now()

	//accesing the paritcular values from struct
	fmt.Println("my order id of order1: ", order1.id)
	fmt.Println("my order status of order1: ", order1.status)

	fmt.Println("my order 1:- ", order1)
	fmt.Println("my order 2:- ", order2)

	// method attached to struct
	changesStatus("confirmed")
	fmt.Print(order1)

}
