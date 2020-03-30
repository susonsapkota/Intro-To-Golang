//Hands-on exercise #6
//Create a program that shows the “if statement” in action.

package main

import "fmt"

func main() {
	num := 8
	if num%2 == 0 {
		fmt.Println("Number is odd")
	} else {
		fmt.Println("Number is even")
	}
}
