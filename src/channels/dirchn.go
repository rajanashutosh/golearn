package main

import "fmt"

func main() {
	//c := make(chan int)
	cr := make(<-chan int)
	cs := make(chan<- int)
	go func() {
		cs <-
			<-cr
	}()

	fmt.Println(cr)
}
