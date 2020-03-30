//Hands-on exercise #1
//Start with this code.
//Instead of using the blank identifier, make sure the code is checking and handling the error.
//https://play.golang.org/p/3W69TH4nON
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	First   string
	Last    string
	Sayings []string
}

func main() {
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	bs, err := json.Marshal(p1)
	if err != nil {
		log.Fatalln("Error while marshalling the json, ", err)
		return
	}
	fmt.Println(string(bs))

}
