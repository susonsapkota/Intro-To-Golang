//Hands-on exercise #3
//Using goroutines, create an incrementer program
//have a variable to hold the incrementer value
//launch a bunch of goroutines
//each goroutine should
//read the incrementer value
//store it in a new variable
//yield the processor with runtime.Gosched()
//increment the new variable
//write the value in the new variable back to the incrementer variable
//use waitgroups to wait for all of your goroutines to finish
//the above will create a race condition.
//Prove that it is a race condition by using the -race flag
//if you need help, here is a hint:  https://play.golang.org/p/FYGoflKQej
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

	for i := 0; i < gs; i++ {
		go func() {
			v := counter
			runtime.Gosched()
			v++
			counter = v
			wg.Done()
		}()

	}
	wg.Wait()
	fmt.Println("count:", counter)
}
