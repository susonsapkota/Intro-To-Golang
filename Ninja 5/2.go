//Hands-on exercise #2
//Take the code from the previous exercise, then store
//the values of type person in a map with the key of last name.
//Access each value in the map. Print out the values, ranging over the slice.
package main

import "fmt"

type person struct {
	firstName       string
	lastName        string
	favouriteFlavor []string
}

func main() {

	p1 := person{
		firstName:       "Ram",
		lastName:        "Lama",
		favouriteFlavor: []string{"Vanilla", "Chocolate",},
	}

	p2 := person{
		firstName:       "Hari",
		lastName:        "Shrestha",
		favouriteFlavor: []string{"Fried", "Chocolate",},
	}

	personMap := map[string]person{
		p1.firstName: p1,
		p2.firstName: p2,
	}
	fmt.Println("Fav flavor of ", p1.firstName)
	for _, v := range p1.favouriteFlavor {
		fmt.Println("\t", v)
	}
	fmt.Println("Fav flavor of ", p2.firstName)
	for _, v := range p2.favouriteFlavor {
		fmt.Println("\t", v)
	}

	for _, singlePerson := range personMap {
		fmt.Println(singlePerson.firstName)
		fmt.Println(singlePerson.lastName)
		fmt.Println("Favourite Flavour :")
		for _, v := range singlePerson.favouriteFlavor {
			fmt.Println(v)
		}
	}
}
