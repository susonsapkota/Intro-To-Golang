//Create your own type “person” which will have an underlying type of “struct”
//so that it can store the following data:
//	first name
//	last name
//	favorite ice cream flavors
//Create two VALUES of TYPE person. Print out the values, ranging over the elements
//in the slice which stores the favorite flavors.
package main

import "fmt"

type person struct {
	firstName      string
	lastName       string
	favoriteFlavor []string
}

func main() {

	p1 := person{
		firstName:      "Ram",
		lastName:       "Lama",
		favoriteFlavor: []string{"Vanilla", "Chocolate"},
	}

	p2 := person{
		firstName:      "Hari",
		lastName:       "Shrestha",
		favoriteFlavor: []string{"Fried", "Chocolate"},
	}

	fmt.Println("Fav flavor of ", p1.firstName)
	for _, v := range p1.favoriteFlavor {
		fmt.Println("\t", v)
	}
	fmt.Println("Fav flavor of ", p2.firstName)
	for _, v := range p2.favoriteFlavor {
		fmt.Println("\t", v)
	}

}
