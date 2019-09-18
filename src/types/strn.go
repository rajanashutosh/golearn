package main

import "fmt"

func strDemo() {

}

const ab = 42
const (
	num1 = iota
	num2
	num3
)

func main() {
	str := "Hello GoLang"
	fmt.Println(str)
	fmt.Printf("%T\n", str)

	bs := []byte (str)
	fmt.Println(bs)

	for i := 0; i < len(str); i++ {
		fmt.Printf("%#U ", str[i])
	}

	fmt.Printf("\n A: %v %T", ab, ab)
	fmt.Printf("\nNumbers: %v %T", num1, num1)
	fmt.Printf("\nNumbers: %v %T", num2, num2)
}
