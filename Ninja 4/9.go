//Using the code from the previous example, add a record to your map.
//Now print the map out using the “range” loop

package main

import "fmt"

func main() {

	records := map[string][]string{
		"bond_james":      {"Shaken, not stirred", "Martinis", "Women"},
		`moneypenny_miss`: {`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           {`Being evil`, `Ice cream`, `Sunsets`},
	}
	records["Suson"] = []string{`Being evil`, `Ice cream`, `Sunsets`}

	for k, v := range records {
		fmt.Println("Key: ", k, "value: ", v)
		for i, val := range v {
			fmt.Println("Index : ", i, "value: ", val)
		}

	}
}
