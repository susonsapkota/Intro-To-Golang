//Hands-on exercise #8
//Create a func which returns a func
//assign the returned func to a variable
//call the returned func
package main

import "fmt"

func returnFunc() func() int {
	return func() int {
		return 509
	}

}
func main() {

	fmt.Println(returnFunc()()) // second parenthesis calling the inside function

}
