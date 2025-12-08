package main

import "fmt"

// func printSlice(items []int){
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

// func printStringSlice(items []string){
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

// func printSlice[T any](items []T){ // we can also use -> [T inerface{}]
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

// func printSlice[T int | string](items []T){
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

func printSlice[T comparable](items []T){
	for _, item := range items {
		fmt.Println(item)
	}
}

// LIFO
type stack struct {
	elements []int
}

// generics in stack
type stackS[T any] struct {
	elements []T
}

func main(){
	nums := []int{1, 2, 3}

	printSlice(nums)

	names := []string { "go", "lang"}
	// printStringSlice(names)

	printSlice(names)

	myStack :=  stack {
		elements: []int{1, 2, 3},
	}

	myStacks := stackS[string] {
		elements: []string{"golang"},
	}

	fmt.Println(myStack)
	fmt.Println(myStacks)
}
