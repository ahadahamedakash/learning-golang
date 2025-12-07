package main

import "fmt"

func main(){
	age := 17

	if age >= 18 {
		fmt.Println("Adult person")
	} else {
		fmt.Println("Is not an adult person")
	}


	var role = "admin"

	if role == "admin" {
		fmt.Println("This is admin")
	}


	// DECLARE VARIABLE INSIDE THE IF CONSTRUCT
	if age := 15; age < 18 {
		fmt.Println("Age is less then 18")
	}
	
}
