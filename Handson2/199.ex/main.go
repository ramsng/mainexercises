package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type review []user

func (u review) Len() int {
	return len(u)
}
func (u review) Swap(i int, j int) {
	u[i], u[j] = u[j], u[i]
	//	return j,i
}

func (u review) Less(i int, j int) bool {
	return u[i].Age < u[j].Age
}

func (u user) String() string {
	return fmt.Sprintln("\nName:", u.First, u.Last, "\nAge:", u.Age, "\nSayings:")
}
func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	//users := []user{u1, u2, u3}

	//	fmt.Println(users)
	//	sort.Sort(review(users))
	//	fmt.Println(users)
	fmt.Println(u1)
	sort.Strings(u1.Sayings)
	for _, b := range u1.Sayings {
		fmt.Println(b)
	}
	fmt.Println(u2)
	sort.Strings(u2.Sayings)
	for _, b := range u2.Sayings {
		fmt.Println(b)
	}
	fmt.Println(u3)
	sort.Strings(u3.Sayings)
	for _, b := range u2.Sayings {
		fmt.Println(b)
	}
}
