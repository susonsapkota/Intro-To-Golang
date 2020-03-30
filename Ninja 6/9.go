//Hands-on exercise #9
//A “callback” is when we pass a func into a func as an argument. For this exercise,
//pass a func into a func as an argument

package main

import "fmt"

func sum(a int, b int) {
	fmt.Println("SUM is", a+b)
}

// this function takes another function as parameter
func callingFunctionInFunction(anotherFunc func(a int, b int), first int, second int) {
	// calling that function with other which in this case is above function named "sum"
	anotherFunc(first, second)

}

func main() {
	// function calling with function as a param
	callingFunctionInFunction(sum, 5, 4)
}
