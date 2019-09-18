package main

import "fmt"

const (
	val int = 4
)

func constMethod() {
	fmt.Printf("\n%v %b", val, val)
	shi := val<<1
	fmt.Printf("\n%v %b", shi, shi)
}

func main() {
	constMethod()
}
