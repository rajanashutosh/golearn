package main

import "fmt"

func array() {
	var x [5]int
	fmt.Println(x)
	x[2] = 1
	fmt.Println(x)
}

func slices() {
	//x := type {values}
	x := []int{1, 2, 4}
	fmt.Println(x)

	for i := 0; i < len(x); i++ {
		fmt.Println(x[i])
	}

	for i, v := range x {
		fmt.Println(i, " ", v)
	}
}

func slicingASlice() {
	x := []int{1, 2, 4, 12, 23, 11, 10}
	fmt.Println(x)
	fmt.Println(x[1:])
	fmt.Println(x[2:4])
	fmt.Println(x[:2])
}

func appendSlice() {
	x := []int{1, 2, 4, 12, 23, 11, 10}
	x = append(x, 0, 9)
	fmt.Println(x)
	y := []int{100, 200}
	x = append(x, y...)
	fmt.Println(x)
}

func deleteFromSlice() {
	x := []int{1, 2, 4, 12, 23, 11, 10}
	x = append(x[2:4])
	fmt.Println(x)
}

func makeSlice() {
	x := make([]int, 10)
	fmt.Println(x)
	x[1] = 2
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
}

func multiDimSlice() {
	str := make([]string, 3)
	str[0] = "test1"
	str[1] = "test2"
	str[2] = "test3"
	str1 := make([]string, 3)
	str1[0] = "row1"
	str1[1] = "row2"
	str1[2] = "row3"
	strArr := [][]string{str, str1}
	fmt.Println(strArr)
}

func main() {
	array()
	slices()
	slicingASlice()
	appendSlice()
	deleteFromSlice()
	makeSlice()
	multiDimSlice()

}
