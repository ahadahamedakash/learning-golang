package main

import "fmt"

func main(){
	// While loop using for
	i:= 1
	for i <= 3 {
		fmt.Println(i)

		i += 1
	}

	// For loop
	for i := 0; i < 3; i++ {
		fmt.Println(i)
		// break
		// continue
	}

	// 1.22 range
	for i := range 3 {
		fmt.Println(i)
	}
}
