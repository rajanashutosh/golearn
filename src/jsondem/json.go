package main

import (
	"encoding/json"
	"fmt"
	"sort"
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
	e3 := Employ{
		FirstName: "Joseph",
		LastName:  "R",
		Age:       43,
		Address: Address{
			AddressLine1: "G205",
			AddressLine2: "Central Park",
			PinCode:      "600119",
		},
	}
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
		Age:       23,
		Address: Address{
			AddressLine1: "Cherukattil",
			AddressLine2: "House",
			PinCode:      "686609",
		},
	}
	ems := [] Employ{e3, e1, e2}
	return ems
}

type ByPersonFirstName [] Employ

func (b ByPersonFirstName) Len() int           { return len(b) }
func (b ByPersonFirstName) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
func (b ByPersonFirstName) Less(i, j int) bool { return b[i].FirstName < b[j].FirstName }

type ByPersonAge []Employ

func (a ByPersonAge) Len() int           { return len(a) }
func (a ByPersonAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByPersonAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

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

	fmt.Println("\nBefore Sort: ", newEmployees)
	sort.Sort(ByPersonFirstName(newEmployees))
	fmt.Println("\nAfter Sort: ", newEmployees)

	fmt.Println("\nBefore Sort: ", newEmployees)
	sort.Sort(ByPersonAge(newEmployees))
	fmt.Println("\nAfter Sort: ", newEmployees)

}
