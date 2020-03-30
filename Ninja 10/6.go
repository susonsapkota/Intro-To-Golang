//Hands-on exercise #6
//write a program that
//puts 100 numbers to a channel
//pull the numbers off the channel and print them
package main

import "fmt"

func main() {
	c := make(chan int)
	quit := make(chan bool)
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
		quit <- true
	}()
	receive(c, quit)

}

func receive(val chan int, quit chan bool) {
	for {
		select {
		case v := <-val:
			fmt.Println(v)
		case <-quit:
			fmt.Println("Program Exited")
			return
		}
	}

}
