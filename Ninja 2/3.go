//Hands-on exercise #3
//Create TYPED and UNTYPED constants. Print the values of the constants.

package main

import "fmt"

const a = 10
const b string = "10"

func main() {
	fmt.Printf("%T\t",a)
	fmt.Println("Untyped: ",a)

	fmt.Printf("%T\t", b)
	fmt.Println("Typed: ",b)
}
