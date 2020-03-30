//Hands-on exercise #10
//Closure is when we have “enclosed” the scope of a variable in some code block.
//For this hands-on exercise, create a func which “encloses” the scope of a variable:

package main

import "fmt"

var globalVar string = "Global" // global var has global scope

func scopeOfVariable() {
	functionVar := "Inside of function " // this variable has scope inside of the function
	print(functionVar)

}
func main() {
	localVar := "Iniside main" // local var has scope inside the main function

	for i := 0; i < 3; i++ {
		fmt.Println(i) // i has the scope inside of this for loop
	}
	fmt.Println(globalVar)
	fmt.Println(localVar)
	scopeOfVariable()

	a := 15
	fmt.Printf("%T\n", a)
	y := &a
	fmt.Printf("%T\n", y)
	z := *y
	fmt.Printf("%T\n", z)

}
