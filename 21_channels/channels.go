package main

import "fmt"

// send
// func processNum(numChan chan int) {
// 	// fmt.Println("processing number", <- numChan)

// 	for num := range numChan {
// 		fmt.Println("processing number", num)

// 		time.Sleep(time.Second * 1)
// 	}
// }

// receive
// func sum(result chan int, num1 int, num2 int) {
// 	res := num1 + num2

// 	result <- res // blocking
// }

// func tast(done chan bool) {
// 	defer func () { done <- true}()

// 	fmt.Println("processing...")
// }

// func emailSender(emailChan chan string, done chan bool) { // (emailChan <- chan string, done chan<- bool) <= by this, the emailChan will be a receive only and done will be send only channel
// 	defer func() { done <- true}()

// 	for email := range emailChan {
// 		fmt.Println("Sending email to", email)

// 		time.Sleep(time.Second)
// 	}
// }

// goroutine synchronizer
func main() {

	// multiple channels
	chanOne := make(chan int)
	chanTwo := make(chan string)

	go func() {
		chanOne <- 10
	}()

	go func() {
		chanTwo <- "pingpong"
	}()

	for i := 0; i < 2; i++ {
		select {
		case chanOneVal := <-chanOne:
			fmt.Println("Received data from chan one", chanOneVal)
		case chanTwoVal := <-chanTwo:
			fmt.Println("Received data from chan two", chanTwoVal)
		}
	}

	/*
		// buffered channel
			emailChan := make(chan string, 100)

			done := make(chan bool)

			go emailSender(emailChan, done)

			for i := 0; i <5; i++ {
				emailChan <- fmt.Sprintf("%d@gmail.com", i)
			}

			fmt.Println("Done sending.")

			// without closing program will go to deadlock
			close(emailChan)

			// emailChan <- "1@example.com"
			// emailChan <- "2@example.com"


			// fmt.Println(<- emailChan)
			// fmt.Println(<- emailChan)

			<-done
	*/

	/*
		// unbuffered channel
			done := make(chan bool)

			go tast(done)

			<- done // block
	*/

	// we can do same task using chanel instead of using waitgroup (multiple)

	/*
		result := make(chan int)

		go sum(result, 4, 5)

		res := <- result

		fmt.Println(res)
	*/

	/*
		numChan := make(chan int)

		go processNum(numChan)

		// numChan <- 5

		for {
			numChan <- rand.IntN(100)
		}
	*/

	// time.Sleep(time.Second * 2)

	// messageChannel := make(chan string)

	// messageChannel <- "ping" // blocking

	// msg := <- messageChannel

	// fmt.Println(msg)
	// // fatal error: all goroutines are asleep - deadlock!
}
