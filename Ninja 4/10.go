//Hands-on exercise #10
//Using the code from the previous example, delete a record from your map.
//Now print the map out using the “range” loop

package main

import "fmt"

func main() {
	records := map[string][]string{
		"bond_james":      {"Shaken, not stirred", "Martinis", "Women"},
		`moneypenny_miss`: {`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           {`Being evil`, `Ice cream`, `Sunsets`},
	}
	fmt.Println("Before Deletion: ")
	for k, v := range records {
		fmt.Println("Key: ", k, "value: ", v)
		for i, val := range v {
			fmt.Println("Index : ", i, "value: ", val)
		}

	}

	delete(records, "no_dr")
	fmt.Println("After Deletion: ")
	for k, v := range records {
		fmt.Println("Key: ", k, "value: ", v)
		for i, val := range v {
			fmt.Println("Index : ", i, "value: ", val)
		}

	}

}
