package main

import (
	"fmt"
	"slices"
)

// SLICE -> DYNAMIC ARRAY
// ALSO HAS USEFUL METHODS
func main(){
	// UNINITIALIZED SLICE IS NIL
	var nums []int

	fmt.Println(nums == nil) // TRUE
	fmt.Println(nums)

	fmt.Println(len(nums)) // 0

	var nums_2 = make([]int, 2)

	fmt.Println(nums_2) // [0 0]
	fmt.Println(cap(nums_2)) // 2 // (capacity) -> cap

	var nums_3 = make([]int, 2, 5) // 5 is the initial capacity

	nums_3 = append(nums_3, 1)
	nums_3 = append(nums_3, 2)
	nums_3 = append(nums_3, 3)

	fmt.Println(nums_3) // [ 0 0 1 2 3]
	fmt.Println(cap(nums_3)) // 5

	var nums_4 = make([]int, 2, 5)

	nums_4[0] = 1 // INSER BY INDEX
	nums_4[1] = 10
	fmt.Println(nums_4) // [1 10]

	// COPY FUNCTION
	var n1 = make([]int, 0, 5)
	n1 = append(n1, 2)

	var n2 = make([]int, len(n1))


	copy(n2, n1)

	fmt.Println(n1, n2)

	// SLICE OPERATOR
	var n3 = []int{ 1, 2, 3}

	fmt.Println(n3[0:2]) // WILL PRINT 0 INDEX TO <2 INDEX
	fmt.Println(n3[0:1]) // will print only 1 index
	fmt.Println(n3[:1]) // will print only 1 index

	fmt.Println(n3[1:]) // WILL NOT EXCLUE => [2 3]

	// SLICE PACKAGE
	var n4 = []int { 1, 2}
	var n5 = []int{ 1, 2}

	fmt.Println(slices.Equal(n4, n5)) // true

	// 2D SLICES
	var nu1 = [][]int{{1,2,3}, {4,5,6}}
	fmt.Println(nu1) // [[1 2 3] [4 5 6]]


}
