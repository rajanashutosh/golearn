package main

import "fmt"

//Declare and assign value to variable x with default value of type integer
var x =10

var y int32 =11

func main() {
	fmt.Println("Identifiers in GoLang")
	num := 10
	fmt.Println("Value of num: ", num)
	sum := sum()
	fmt.Println("Sum is: ", sum)
	vardemo()
	fmt.Println("Global X: ", x)
	fmt.Println("Global Y of int32 type: ",y)
}

func sum() int {
	//Declaring and assigning value using short declaration operation
	num := 21
	sum := 100 + num
	return sum
}

func vardemo() {
	var num = 10
	fmt.Println("Var keyword: ", num)
}
