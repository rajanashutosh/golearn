package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("OS: \t", runtime.GOOS)
	fmt.Println("Arch: \t", runtime.GOARCH)
	fmt.Println("CPU: \t", runtime.NumCPU())
	fmt.Println("GORoutine: \t", runtime.NumGoroutine())
	fmt.Println("GoCall: \t", runtime.NumCgoCall())
	wg.Add(1)
	go foo()
	wg.Add(1)
	fmt.Println("CPU: \t", runtime.NumCPU())
	fmt.Println("GORoutine: \t", runtime.NumGoroutine())
	fmt.Println("GoCall: \t", runtime.NumCgoCall())
	go bar()
	fmt.Println("CPU: \t", runtime.NumCPU())
	fmt.Println("GORoutine: \t", runtime.NumGoroutine())
	fmt.Println("GoCall: \t", runtime.NumCgoCall())
	wg.Wait()

}

func foo() {
	for i := 1; i < 10; i++ {
		fmt.Println("Foo: ", i)
	}
	wg.Done()
}

func bar() {
	for i := 1; i < 10; i++ {
		fmt.Println("Bar:", i)
	}
	wg.Done()
}
