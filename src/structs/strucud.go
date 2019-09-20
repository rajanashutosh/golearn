package main

import "fmt"

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

func (p person) setPersonInfo(){
	fmt.Println(p)
}

func setPersonInfo() {
	p1 := person{firstName: "Ashutosh", lastName: "Rajan", age: 30, address: address{addressLine1: "CPW", addressLine2: "SLN", pinCode: "600119"},}
	fmt.Println(p1)
}

func main() {
	//setPersonInfo()
	p1 := person{firstName: "Ashutosh", lastName: "Rajan", age: 30, address: address{addressLine1: "CPW", addressLine2: "SLN", pinCode: "600119"},}
	p1.setPersonInfo();

}
