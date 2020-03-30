//Hands-on exercise #4
//Fix the race condition you created in the previous exercise by using a mutex
//it makes sense to remove runtime.Gosched()
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	counter := 0

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	var mu sync.Mutex
	for i := 0; i < gs; i++ {
		go func() {
			v := counter
			v++
			counter = v
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("count:", counter)
}
