package main

import "fmt"

// type MyType string // custom type
// -> enums created by type and const
type OrderStatus int
// type OrderStatus2 string

const (
	Received OrderStatus = iota // iota is a predeclared identifier representing the untyped integer ordinal number of the current const specification in a (usually parenthesized) const declaration. It is zero-indexed.
	Confirmed
	Prepared
	Delivered
)

// const (
// 	Received OrderStatus2 = "received"
// 	Confirmed = "confirmed"
// 	Prepared = "prepared"
// 	Delivered = "delivered"
// )

func changeOrderStatus(status OrderStatus){
	fmt.Println("change order status: ", status)
}

// enumerated types
func main(){
	changeOrderStatus(Received) // 0
	changeOrderStatus(Delivered) // 3
	// changeOrderStatus("prepared")
}
