package main

import "fmt"

func main() {

	// for init; condition; increment/decrement
	for i := 0; i < 10; i++ {
		fmt.Print(" ", i)
	}
	fmt.Println()
	for i := 0; i < 10; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print(j, " ")
		}
		fmt.Println()
	}

	vla := 0
	for vla < 20 {
		vla++
		fmt.Println(vla)
	}

	vla = 0
	for {
		if vla > 9 {
			break
		}
		fmt.Println(vla)
		vla++
	}

	x := 0
	for x < 10 {
		x++
		if x == 1 {
			fmt.Println("Value is 1")
		} else if x == 2 {
			fmt.Println("Value is 2")
		} else if x == 3 {
			fmt.Println("Value is 3")
		} else if x == 4 {
			fmt.Println("Value is 4")
		} else {
			fmt.Println("Value is not")
		}
	}

}
