package main

import "fmt"

// NUMBER SEQUENCE OF SPECIFIC LENGTH
func main(){
	// ADD ZEROS
	var nums [4]int

	// ARRAY LENGTH
	// fmt.Println(len(nums))

	nums[0] = 1
	fmt.Println(nums) // [1 0 0 0]

	// ADDED FASLSE TO ALL
	var vals [4]bool
	vals[0] =  true
	fmt.Println(vals) // [true false false false]

	// ADDED EMPTY STRINGS
	var name [4]string
	name[0] =  "Go"
	fmt.Println(name) // [Go   ]

	// TO DECLARE IN SINGLE LINE
	nums_2 := [3]int{ 1, 2, 3}
	fmt.Println(nums_2)

	// 2D ARRAY 
	nums_3 := [2][2]int{{1,2 },{3, 4}}
	fmt.Println(nums_3)

	// - fixed size, that is predictable
	// - Memory optimization
	// - Constant time access
	// (slices used maximum time in go)
}
