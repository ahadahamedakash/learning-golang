package main

import "fmt"

const age = 20

func main() {
	const name string = "Golang"

	fmt.Println(age)

	const (
		post = 5000
		host = "localhost"
	)

	fmt.Println(post, host)

}