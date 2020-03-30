//Hands-on exercise #1
//get this code working: https://play.golang.org/p/j-EA6003P0
//using func literal, aka, anonymous self-executing func

package main

import (
	"fmt"
)

// using function literals
//func main() {
//	c := make(chan int)
//
//	go func() {
//		c <- 42
//	}()
//
//	fmt.Println(<-c)
//
//}

// using buffered channels

func main() {
	c := make(chan int, 1)

	c <- 42

	fmt.Println(<-c)

}
