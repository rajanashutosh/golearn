package main

import "fmt"

func main() {
	ch := make(chan int)
	//Send data to channel
	go bar(ch)

	//Receive from channel
	foo(ch)

	fmt.Println("Send and receive completed")

}

//Receive from channel
func foo(ch <-chan int) {
	fmt.Println(<-ch)
}

//Send from channel
func bar(ch chan<- int) {
	ch <- 100
}
