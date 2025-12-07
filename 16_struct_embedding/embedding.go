package main

import "fmt"

type customer struct {
	name string
	phone string
}

type order struct {
	id string
	amount float32
	quantity int
	customer
}

func main(){
	// newCustomer := customer{
	// 	name: "Jhon Doe",
	// 	phone: "123456789",
	// }

	newOrder := order {
		id: "1",
		amount: 30.99,
		quantity: 10,
		// customer: newCustomer,
		// or
		customer: customer{
			name: "Jhon",
			phone: "987654321",
		},
	}

	fmt.Println(newOrder) // {1 30.99 10 {Jhon Cena 987654321}}

	newOrder.customer.name = "Jhon Cena"

	fmt.Println(newOrder) // {1 30.99 10 {Jhon Doe 123456789}}

}