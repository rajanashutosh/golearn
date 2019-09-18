package main

import "fmt"

const (
	num = iota
	kb = 1 << (iota *10)
	mb = 1 << (iota *10)
	gb = 1 << (iota *10)
)

func main(){
	//fmt.Printf("\n%v %b", num, num)
	fmt.Printf("\n%v %b", kb, kb)
	fmt.Printf("\n%v %b", mb, mb)
	fmt.Printf("\n%v %b", gb, gb)
}
