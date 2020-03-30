//Using the code from the previous example, use SLICING to create the following new slices which are then printed:
//[42 43 44 45 46]
//[47 48 49 50 51]
//[44 45 46 47 48]
//[43 44 45 46 47]

package main

import "fmt"

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	firstFive := x[:5]
	secondFive := x[5:]
	thirdFive := x[2:7]
	fourthFive := x[1:6]
	fmt.Println(firstFive)
	fmt.Println(secondFive)
	fmt.Println(thirdFive)
	fmt.Println(fourthFive)

}
