//Hands-on exercise #3
//Starting with this code, pull the values off the channel using a for range loop
//https://play.golang.org/p/sfyu4Is3FG

package main

import (
	"fmt"
)
func receive(ch <-chan int, quit chan bool) {
	for v := range ch {
		fmt.Println(v)
	}
}

func main() {
	c := gen()
	receive(c, nil)

	fmt.Println("about to exit")
}

func gen() <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()
	return c
}

