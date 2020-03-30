//Hands-on exercise #3
//Use the “defer” keyword to show that a deferred func runs after the func containing it exits.
package main

import "fmt"

func main() {
	defer printFirst() // this comes after completing last statement but before function ends
	printSecond()
}

func printFirst() {
	fmt.Println("This is from printFirst function")
}

func printSecond() {
	fmt.Println("This is from printSecond function")
}
