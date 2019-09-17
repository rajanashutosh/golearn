package main

import "fmt"

//GoLang is statically typed language and variable can store values of types which are assigned
//to it or declared of type
var a int = 10

type myType int


//User defined type
var m myType

func main() {
	var x = 10
	var str = "hello"
	fmt.Println("Type declaration")

	fmt.Println("X:= ", x)
	fmt.Printf("X of type %T ", x)

	fmt.Println("Var:= ", str)
	fmt.Printf("Var of type %T \n", str)

	s := fmt.Sprintf("SPrintf Sample %v", x)
	fmt.Println(s)
	fmt.Println("A:=", a)
	m = 100
	//Type Conversion
	a = int(m)
	fmt.Println("M:=", m)
	fmt.Println("A:=", a)
}
