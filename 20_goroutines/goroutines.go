package main

import (
	"fmt"
	"sync"
)

// func task (id int){
// 	fmt.Println("Doing tast", id)
// }

func task (id int, w *sync.WaitGroup){
	defer w.Done() // defer runs after functions execution

	fmt.Println("Doing tast", id)
}

func main(){
	// waitgroud
	var wg sync.WaitGroup

	// for i:= 0; i <= 10; i++ {
	// 	// go func (i int) { // its a good practice to receive input though it works without taking
	// 	// 	fmt.Println(i)
	// 	// } (i)

	// 	go task(i) // runs concurrently
	// }


	for i:= 0; i <= 10; i++ {
		// go func (i int) { // its a good practice to receive input though it works without taking
		// 	fmt.Println(i)
		// } (i)

		wg.Add(1)
		go task(i, &wg) // runs concurrently
	}

	// wait group -> add, run, wait
	wg.Wait()

	// without blocking we can see our result
	// time.Sleep(time.Second * 2)
}
