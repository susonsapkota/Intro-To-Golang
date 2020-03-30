//Hands-on exercise #1
//in addition to the main goroutine, launch two additional goroutines
//each additional goroutine should print something out
//use waitgroups to make sure each goroutine finishes before your program exists

package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func routineFirst() {
	fmt.Println("Printing from first")
	time.Sleep(4 * time.Second)
	fmt.Println("Finished from first")
	wg.Done()
}
func routineSecond() {
	fmt.Println("Printing from Second")
	time.Sleep(5 * time.Second)
	fmt.Println("Finished from Second")
	wg.Done()

}
func main() {
	fmt.Println("No of goroutines before: ", runtime.NumGoroutine())
	wg.Add(2)
	go routineFirst()
	go routineSecond()
	fmt.Println("executes all task on main waiting for rest of the task to complete.")
	fmt.Println("No of goroutines running: ", runtime.NumGoroutine())
	wg.Wait()
	fmt.Println("all task complete.")
}
