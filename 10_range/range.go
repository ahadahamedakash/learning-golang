package main

import "fmt"

func main(){
	nums := []int {6, 7, 8}

	// for i:=0; i <len(nums); i++{
	// 	fmt.Println(nums[i])
	// }

	sum := 0

	for _, num := range nums {
		sum += num
		// fmt.Println(num)
	}

	// fmt.Println(sum)

	m := map[string]string { "firstName": "Jhon", "lastName": "Doe"}

	for k, v := range m {
		fmt.Println(k, v)
		/*
			OUTPUT: 
				firstName Jhon
				lastName Doe
		*/

		// k = keys, v = values (convention)
	}

	// UNICODE CODE POINT RUNE
	// STARTING BYTE OF RUNE
	// 255 -> 1 BYTE , 2 BYTE
	for i, c := range "golang" {
		// fmt.Println(i, c)
		fmt.Println(i, string(c))
	}
}