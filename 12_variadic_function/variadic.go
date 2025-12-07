package main

import "fmt"

func sum(nums ...int) int{ // we get slice here
	// func sum(nums ...interfaces{}) int{} // by the interfaces we can send anything
	total := 0
	for _, i:= range nums {
		total += i
	}

	return total
}

func main(){
	// fmt.Println(1, 2, 3, 4, 5)

	grandTotal := sum(1, 2, 3, 4, 5, 6, 7,8 , 9, 10) // we can send as much parameter as we want

	fmt.Println(grandTotal)

	nums:= []int{3, 4, 5, 6}
	
	fmt.Println(sum(nums...))

}