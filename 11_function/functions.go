package main

import "fmt"

// func add(a, b int) int {}  -> if both are same type, we can give only one
func add(a int, b int) int {
	return a + b
}

func getLanguage()(string, string, string){
	return "golang", "js", "c"
}

func processIt(fn func(a int) int){
	fn(1)
}

func processItAgain() func(a int) int {
	return func(a int)int {
		return 4
	}
}

func main(){
	result := add(3, 6)

	fmt.Println(result)

	fmt.Println(getLanguage()) // fmt.Println(lang1, lang2, lang3)
	lang1, lang2, lang3 := getLanguage()

	fmt.Println(lang1) // golang
	fmt.Println(lang2) // js
	fmt.Println(lang3) // c

	fn:= func(a int) int {
		return 2
	}

	processIt(fn)

	fn2 := processItAgain() // returning a function
	fn2(4)

}