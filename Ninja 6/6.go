//Hands-on exercise #6
//Build and use an anonymous func

package main

import "fmt"

func main() {

	func(a int, b int) {
		fmt.Println("Sum is :", a+b)
	}(5, 4)
}
