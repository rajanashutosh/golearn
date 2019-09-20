package main

import "fmt"

func main() {
	func() {
		fmt.Println("Anony")
	}()

	f := func() {
		fmt.Println("First Class Citizen")
	}
	f()

	x := foo()
	fmt.Println(x())
}

func foo() func() int {
	return func() int {
		return 10
	}
}
