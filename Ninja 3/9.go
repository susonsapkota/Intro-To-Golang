//Hands-on exercise #9
//Create a program that uses a switch statement with the switch expression specified as a variable of TYPE string with the IDENTIFIER “favSport”.

package main

import "fmt"

func main() {
	favSport := "Football"
	switch favSport {
	case "Football":
		fmt.Println("Football is awesome")
	case "Basketball":
		fmt.Println("Basketball is awesome")
	case "Cricket":
		fmt.Println("Cricket is awesome")
	case "Volleyball":
		fmt.Println("Volleyball is awesome")
	}
}
