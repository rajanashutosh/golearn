package main

import "fmt"

func foo1() {
	fmt.Println("Foo ")
}

func incrementor() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}
func main() {
	foo1()
	a:= incrementor()
	b:= incrementor()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(b())

}
