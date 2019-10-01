package main

import "fmt"

func main() {
	ch := make(chan int) //channel creation

	go func() { // (goroutine) anon method to push value into channel
		ch <- 43
	}()

	fmt.Println(<-ch) // read from channel and syncs with goroutine
}
