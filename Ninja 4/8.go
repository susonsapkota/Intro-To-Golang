//Create a map with a key of TYPE string which is a person’s “last_first” name, and a value of TYPE
//[]string which stores their favorite things. Store three records in your map. Print out all of the
//values, along with their index position in the slice.
//
//`bond_james`, `Shaken, not stirred`, `Martinis`, `Women`
//`moneypenny_miss`, `James Bond`, `Literature`, `Computer Science`
//`no_dr`, `Being evil`, `Ice cream`, `Sunsets`

package main

import "fmt"

func main() {
	records := map[string][]string{
		"bond_james":      {"Shaken, not stirred", "Martinis", "Women"},
		`moneypenny_miss`: {`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           {`Being evil`, `Ice cream`, `Sunsets`},
	}
	for k, v := range records {
		fmt.Println("Key: ", k, "value: ", v)
		for i, val := range v {
			fmt.Println("Index : ", i, "value: ", val)
		}

	}
}
