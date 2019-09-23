package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	var counter int64 = 0
	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)
	fmt.Println("GORoutine: \t", runtime.NumGoroutine())
	for i := 0; i < gs; i++ {
		go func() {
			//v := counter
			atomic.AddInt64(&counter, 1)
			runtime.Gosched()
			fmt.Print("Counter: ", atomic.LoadInt64(&counter))
			//v++
			//counter = v
			wg.Done()
		}()
		fmt.Println("GORoutine: \t", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("GORoutine: \t", runtime.NumGoroutine())
	fmt.Println("Counter: ", counter)
}
