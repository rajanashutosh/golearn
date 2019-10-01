package main

import "fmt"

func main() {
	eve := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	//Send data to channel
	go send(eve, odd, quit)

	//Receive data from channel
	receive(eve, odd, quit)

	fmt.Println("About to exit...")
}

func send(eve, odd, quit chan<- int) {
	for i := 0; i <= 100; i++ {
		if i%2 == 0 {
			eve <- i
		} else {
			odd <- i
		}
	}
	quit <- 0
}

func receive(eve, odd, quit <-chan int) {
	for {
		select {
		case v := <-eve:
			fmt.Println(v)
		case v := <-odd:
			fmt.Println(v)
		case v := <-quit:
			fmt.Println(v)
			return
		}
	}

}
