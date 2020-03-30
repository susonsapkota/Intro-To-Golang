//Hands-on exercise #4
//Write a program that:
//	assigns an int to a variable
//	prints that int in decimal, binary, and hex
//	shifts the bits of that int over 1 position to the left, and assigns that to a variable
//	prints that variable in decimal, binary, and hex

package main

import "fmt"

func main() {
	intVar := 10
	fmt.Printf("%d\n", intVar)
	fmt.Printf("%b\n", intVar)
	fmt.Printf("%#x\n\n", intVar)

	shiftedInt := intVar << 1

	fmt.Printf("%d\n", shiftedInt)
	fmt.Printf("%b\n", shiftedInt)
	fmt.Printf("%#x\n", shiftedInt)
}
