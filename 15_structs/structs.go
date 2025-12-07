package main

import (
	"fmt"
	"time"
)

// ORDER STRUCTS
type order struct {
	id string
	amount float32
	status string
	createdAt time.Time // NANOSECOND PRECISION
}

func newOrder (id string, amount float32, status string) *order {
	// initial setup goes here
	myOrder := order {
		id: id,
		amount: amount,
		status: status,
	}

	return &myOrder
}

func (o *order) changeStatus(status string){
	o.status = status
}

func (o *order) getAmount() float32 {
	return o.amount
}

// STRUCTS -> CUSTOM DATA STRUCTURE (LIKE CLASSES)
func main(){
	// var order order = 
	myOrder := order {
		id: "1a2b3c",
		amount: 50.99,
		status: "delivered",
	}

	fmt.Println("Order Struct", myOrder) // Order Struct {1a2b3c 50.99 delivered {0 0 <nil>}}
	myOrder.createdAt = time.Now()

	fmt.Println("Order Struct", myOrder) // Order Struct {1a2b3c 50.99 delivered {13998759018480265923 16382 0x57fe40}}

	myOrder2 := order {
		id: "2b3c4d",
		amount: 100,
		status: "pending",
		createdAt: time.Now(),
	}

	fmt.Println("Order Struct two: ", myOrder2) // Order Struct two:  {2b3c4d 100 pending {13998759438242300813 51568 0x57fe40}}

	myOrder2.changeStatus("delivered")

	fmt.Println("Order Struct two: ", myOrder2) // Order Struct two:  {2b3c4d 100 delivered {13998759438242300813 51568 0x57fe40}}

	fmt.Println(myOrder2.getAmount()) // 100

	/*
		if we don't set any field, default value will be 0
	*/

	//
	myNewOrder := newOrder("12345", 30.99, "received")

	fmt.Println(myNewOrder)

	//
	// SINGLE LINE STRUCT
	language := struct {
		name string
		isGood bool
	} {"Golang", true}

	fmt.Println(language)

}