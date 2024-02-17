package main

import (
	"encoding/json"
	"fmt"
	"log"
)

/*
	type User struct {
		Name string
		Age  int
		Auth bool
	}
*/
type Jsonstruct struct {
	First   string   `json:"First"`
	Last    string   `json:"Last"`
	Age     int      `json:"Age"`
	Sayings []string `json:"Sayings"`
}

func (j Jsonstruct) String() string {
	return fmt.Sprintln("\nName:\t", j.First, j.Last, "\nAge :\t", j.Age, "\nSayings: ", j.Sayings)
}

func main() {
	jsonblog := `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`
	var u1 []Jsonstruct
	err := json.Unmarshal([]byte(jsonblog), &u1)
	if err != nil {
		log.Fatal("error", err)
	}
	fmt.Println("json input :", jsonblog)
	fmt.Println(u1)
}
