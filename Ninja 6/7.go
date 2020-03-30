//Hands-on exercise #7
//Assign a func to a variable, then call that func
package main

import "fmt"

func main() {
	a := func(a int, b int) int {
		fmt.Println("Sum is :", a+b)
		return a + b
	}
	a(1, 3)
}
