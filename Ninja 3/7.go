//Hands-on exercise #7
//Building on the previous hands-on exercise, create a program that uses “else if” and “else”.
package main

import "fmt"

func main() {
	num := 8
	if num%2 == 0 {
		fmt.Println("Number is divisible by 2")
	} else if num%3 == 0 {
		fmt.Println("Number is divisible by 3")
	} else {
		fmt.Println("Number is neither divisible by 3 nor by 2")
	}
}
