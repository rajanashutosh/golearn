package main

import "fmt"

func main() {
	a := 42

	fmt.Printf("%T %v\n", a, a)
	fmt.Printf("%T %v\n", &a, &a)
	b := &a
	fmt.Printf("%T %v\n", b, b)
	fmt.Printf("%T %v\n", *b, *b)
}
