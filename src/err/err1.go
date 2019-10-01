package main

import (
	"fmt"
	"log"
)

func main() {
	log.Println("Execution started")
	n, err := fmt.Println("Hellow")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n)
	log.Println("About to exit...")
}
