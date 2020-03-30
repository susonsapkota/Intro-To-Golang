//Hands-on exercise #6
//Using iota, create 4 constants for the NEXT 4 years. Print the constant values.

package main

import "fmt"

const (
	_ = (2020) + iota
	year21
	year22
	year23
	year24
)

func main() {

	fmt.Println(year21, year22, year23,year24)

}
