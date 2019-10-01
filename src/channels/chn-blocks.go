package main

import "fmt"

func main() {
	ch := make(chan int)

	//Send to channel
	go func() {
		for i := 1; i < 100; i++ {
			ch <- i
		}
		close(ch)
	}()

	//Recieve from channel
	for v := range ch {
		fmt.Println(v)
	}

}
