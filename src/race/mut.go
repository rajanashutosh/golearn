package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var counter int = 0
	const gs = 100
	fmt.Println("Before GORoutine: \t", runtime.NumGoroutine())
	fmt.Println("Counter: ", counter)
	var ws sync.WaitGroup
	var mu sync.Mutex
	ws.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			v := counter
			v++
			runtime.Gosched()
			counter = v
			mu.Unlock()
			ws.Done()
		}()
		fmt.Println("Counter: ", counter)
	}
	ws.Wait()
	fmt.Println("After GORoutine: \t", runtime.NumGoroutine())
	fmt.Println("Counter: ", counter)
}
