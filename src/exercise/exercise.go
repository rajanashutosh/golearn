package main

import "fmt"

func assignValues() {
	x := 42
	y := "James Bond"
	z := true

	fmt.Printf("Value of x:= %v %T \n", x, x)
	fmt.Printf("Value of y:= %v %T \n", y, y)
	fmt.Printf("Value of z:= %v %T \n", z, z)
}

var x int
var y string
var z bool


var a int = 42
var b string = "James Bond"
var c bool = true

type myType int

func main() {
	fmt.Println("HandsOn Exercise")
	fmt.Println("Assign Value method call")
	assignValues()
	fmt.Println("Compiler Assigned Values")
	fmt.Printf("Value of x:= %v %T \n", x, x)
	fmt.Printf("Value of y:= %v %T \n", y, y)
	fmt.Printf("Value of z:= %v %T \n", z, z)

	val:=fmt.Sprintf("\n A:= %v \n B:=%v \n C:= %v ", a,b,c)
	fmt.Println("Value of string: ", val)

	var num myType =10
	fmt.Println("Num: ", num)
	fmt.Printf("Num %v %T \n", num, num)
	var num1 int =11
	fmt.Println("Num1: ", num1)
	fmt.Printf("Num1 %v %T \n", num1, num1)

	num1 = int (num)
	fmt.Println("Num: ", num1)
	fmt.Printf("Num1 %v %T \n", num1, num1)
}
