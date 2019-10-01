package main

import (
	"fmt"
	"sync"
)

func main() {
	even := make(chan int)
	odd := make(chan int)
	fanIn := make(chan int)

	//send data
	go sendData(even, odd)

	//Receive Data
	go recData(even, odd, fanIn)

	for v := range fanIn {
		fmt.Println(v)
	}

	fmt.Println("About to exit..")
}

func sendData(e, o chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	close(e)
	close(o)
}

func recData(e, o <-chan int, f chan<- int) {
	var ws sync.WaitGroup
	ws.Add(2)
	go func() {
		for v := range e {
			f <- v
		}
		ws.Done()
	}()

	go func() {
		for v := range o {
			f <- v
		}
		ws.Done()
	}()
	ws.Wait()
	close(f)
}
