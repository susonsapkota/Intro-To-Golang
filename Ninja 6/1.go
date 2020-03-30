//Review
//	functions
//	purpose of functions
//	abstract code
//	code reusability
//		DRY - Don’t Repeat Yourself
//	func, receiver, identifier, params, returns
//	parameters vs arguments
//	variadic funcs
//	multiple “variadic” params
//	multiple “variadic” args
//	returns
//	multiple returns
//	named returns - yuck!
//	func expressions
//	assigning a func to a variable
//	callbacks
//	passing a func into another func as an argument
//	closure
//	one scope enclosing another
//	variables declared in the outer scope are accessible in inner scopes
//	closure helps us limit the scope of variables
//	recursion
//	a func that calls itself
//	factorial
//	life philosophy
//	focus on what’s important; not upon what’s urgent
//
//
//	Hands on exercise
//create a func with the identifier foo that returns an int
//create a func with the identifier bar that returns an int and a string
//call both funcs
//print out their results

package main

import "fmt"

func foo(a int, b int) int {
	return a + b
}

func bar(a int, b int) (int, string) {
	return a - b, "Subtraction"

}
func main() {
	fmt.Println(foo(1, 2))
	fmt.Println(bar(1, 2))
}
