//Hands-on exercise #2
//https://play.golang.org/p/b_UuCcZag9
//Starting with this code, unmarshal the JSON into a Go data structure.
//Hint: https://mholt.github.io/json-to-go/
package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First   string   `json:"First"`
	Last    string   `json:"Last"`
	Age     int      `json:"Age"`
	Sayings []string `json:"Sayings"`
}

func main() {
	s := `[{"First":"James","Last":"Bond","Age":32,
"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},
{"First":"Miss","Last":"Moneypenny","Age":27,
"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},
{"First":"M","Last":"Hmmmm","Age":54,
"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`

	fmt.Println(s)
	// convert to bytes
	bs := []byte(s)

	// create a variable to store the unmarshaled json
	var people []person
	// first the bytes of string then the location 'address' if the variable
	err := json.Unmarshal(bs, &people)
	// error handling
	if err != nil {
		fmt.Println("Error")
		return
	}
	// displaying the data s
	for i, v := range people {
		fmt.Println(i, v)

	}

}
