//Hands-on exercise #7
//Create a slice of a slice of string ([][]string). Store the following data in the multi-dimensional slice:
//
//"James", "Bond", "Shaken, not stirred"
//"Miss", "Moneypenny", "Helloooooo, James."
//
//Range over the records, then range over the data in each record.

package main

import "fmt"

func main() {

	singleSlice := []string{"James", "Bond", "Shaken, not stirred"}
	nextSlice := []string{"Miss", "Moneypenny", "Helloooooo, James."}
	fmt.Println(singleSlice)
	fmt.Println(nextSlice)
	sliceOfSlice := [][]string{singleSlice, nextSlice}
	fmt.Println(sliceOfSlice)

	for i, single := range sliceOfSlice {
		fmt.Println(i)
		for j, v := range single {
			fmt.Println("Index:", j, " Value: ", v)
		}
	}
}
