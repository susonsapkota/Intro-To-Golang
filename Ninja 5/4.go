//Create and use an anonymous struct.

package main

import (
	"fmt"
)

func main() {

	person := struct {
		name     string
		age      int
		interest []string
		favNum   map[string]string
	}{
		name:     "Suson",
		age:      42,
		interest: []string{"Football", "Videogames", "Movies"},
		favNum:   map[string]string{"home": "01410000", "mobile": "9841617181"},
	}
	fmt.Println(person)

	fmt.Println("Interest:")
	for _, v := range person.interest {
		fmt.Println("\t", v)
	}
	fmt.Println("Phone Numbers:")
	for k, v := range person.favNum {
		fmt.Println("\t", k, v)
	}
}
