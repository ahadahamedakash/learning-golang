package main

import "fmt"

// BY VALUE
// func changeNum(num int){
// 	num = 5
// 	fmt.Println(num) // 5
// }

// func main(){
// 	num := 1
// 	changeNum(num)

// 	fmt.Println("After change in main: ", num) // 1
// }

// BY REFERENCE
func changeNum(num *int){
	*num = 5 // DEREFERENCE
	fmt.Println(num) // MEMORY ADDRESS
	fmt.Println(*num) // VALUE
}

func main(){
	num := 1
	changeNum(&num)

	fmt.Println("MEMORY ADDRESS: ", &num) // MEMORY ADDRESS

	fmt.Println("After change in main: ", num) // 5
}