package main

import (
	"encoding/json"
	"fmt"
)

type Address struct {
	AddressLine1 string
	AddressLine2 string
	PinCode      string
}
type Employ struct {
	FirstName string
	LastName  string
	Age       int
	Address   Address
}

func (employ *Employ) setEmployData() []Employ {
	e1 := Employ{
		FirstName: "Ashutosh",
		LastName:  "Rajan",
		Age:       30,
		Address: Address{
			AddressLine1: "G205",
			AddressLine2: "Central Park",
			PinCode:      "600119",
		},
	}
	e2 := Employ{
		FirstName: "Chinku",
		LastName:  "",
		Age:       33,
		Address: Address{
			AddressLine1: "Cherukattil",
			AddressLine2: "House",
			PinCode:      "686609",
		},
	}
	ems := [] Employ{e1, e2}
	return ems
}

func main() {
	employ := Employ{}
	employees := employ.setEmployData()
	//fmt.Printf("\nEmployees: %v\n", employees)
	bs, err := json.Marshal(employees)
	if err != nil {
		fmt.Println("Marshalling error")
		fmt.Printf("\n%v", err)
	} else {
		jsonStr := string(bs)
		fmt.Println("Marshalling completed")
		fmt.Printf("\n%v", jsonStr)
	}

	newEmployees := []Employ{}

	err = json.Unmarshal(bs, &newEmployees)
	if err != nil {
		fmt.Println("Unmarshalling error")
		fmt.Printf("\n%v", err)
	} else {
		fmt.Println("Unmarshalling completed")
		fmt.Printf("\n%v", newEmployees)
	}

}
