//Hands-on exercise #5
//Create a variable of type string using a raw string literal. Print it.

package main

import "fmt"

func main() {
	stringVar := `This also 
takes accounts	to
returns 						or tabs.
"Just like that"
`
	fmt.Println(stringVar)
}
