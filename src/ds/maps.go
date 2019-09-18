package main

import "fmt"

func mapDemo() {
	m := map[string]int{
		"Nam1":  1,
		"Test1": 2,
	}

	fmt.Println(m)
	fmt.Println(m["Nam1"])

	v, ok := m["Ash"]
	fmt.Println(v)
	fmt.Println(ok)

	if v, ok := m["Nam1"]; ok {
		fmt.Println(v)
	}

	m["Ashu"] = 33
	for k, v := range m {
		fmt.Println(k, v)
	}

	if v, ok := m["Ashu"]; ok {
		fmt.Println(v)
		delete(m, "Ashu")
	}

	for k, v := range m {
		fmt.Println(k, v)
	}
}

func main() {
	mapDemo()

}
