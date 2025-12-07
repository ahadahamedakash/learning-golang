package main

import (
	"fmt"
	"time"
)

func main (){
	i := 2

	switch i {
	case 1: 
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	default:
		fmt.Println("Other")
	}

	// MULTIPLE CONDITION SWTICH
	switch time.Now().Weekday(){
	case time.Friday, time.Saturday:
		fmt.Println("It's weekend")
	default:
		fmt.Println("It's workday")
	}

	// TYPE SWITCH
	whoAmI := func(i interface{}){
		switch t:= i.(type){
		case int: 
		fmt.Println("Its an integer")
		case string:
			fmt.Println("Its a string")
		case bool:
			fmt.Println("Its a boolean")
		default:
			fmt.Println("Other", t)
		}
	}

	whoAmI(true)
	whoAmI("String")
	whoAmI(22)
}
