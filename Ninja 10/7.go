//Hands-on exercise #7
//write a program that
//launches 10 goroutines
//each goroutine adds 10 numbers to a channel
//pull the numbers off the channel and print them
package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		for j := 0; j < 10; j++ {
			for i := 0; i < 10; i++ {
				c <- i
			}
		}
	}()

	for h := 0; h < 100; h++ {
		fmt.Println(h, <-c)
	}
}
