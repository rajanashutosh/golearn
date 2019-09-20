package main

import "fmt"

type human interface {
	speak()
}

type person struct {
	firstName string
	lastName  string
	age       int
	address   address
}

type address struct {
	addressLine1 string
	addressLine2 string
	pinCode      string
}

func (p person) speak() {
	fmt.Println(p)
}

func (a address) speak() {
	fmt.Println(a)
}

func main() {
	p1 := person{firstName: "Ashutosh", lastName: "Rajan", age: 30, address: address{addressLine1: "CPW", addressLine2: "SLN", pinCode: "600119"},}

	address := address{addressLine1: "CPW", addressLine2: "SLN", pinCode: "600119"}
	human(p1).speak()
	human(address).speak()
}
