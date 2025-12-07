package main

import (
	"fmt"
	"maps"
)

// maps(associative data structure) -> hash, objects, dict
func main(){
	// CREATING MAP
	m := make(map[string]string) // sirst string is type of key and second one is value type
	m["name"] = "go"
	m["area"] = "backend"

	// get an element
	fmt.Println(m["area"]) // backend

	m1 := make(map[string]int)
	m1["roll"] = 1
	m1["number"] = 2

	fmt.Println(m1["roll"]) // 1
	fmt.Println(m1["number"]) // 2
	fmt.Println(m1["rol"]) // 0 (doesn't exist)


	// LENGTH
	fmt.Println(len(m1))

	fmt.Println(m1)

	// DELETE A SPECIFIC KEY
	delete(m1, "roll")
	// fmt.Println(m1)

	// DELETE ALL
	clear(m1)


	m2 := map[string]int{ "price": 40, "quantity": 4}

	fmt.Println(m2)

	v, ok := m2["quantity"]

	fmt.Println(v) // 4

	if ok {
		fmt.Println("All okay")
	} else {
		fmt.Println("Not okay")
	}


	m3 := map[string]int{ "price": 40, "quantity": 4}
	m4 := map[string]int{ "price": 40, "quantity": 4}

	// TO CHECK IF BOTH ARE EQUAL OR NOT
	fmt.Println(maps.Equal(m3, m4)) // true

}