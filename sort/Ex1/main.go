package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

func (P Person) String() string {
	return fmt.Sprintln(P.Name, ".", P.Age)
}

type byage []Person

func (x byage) Len() int           { return len(x) }
func (x byage) Less(i, j int) bool { return x[i].Age < x[j].Age }
func (x byage) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

type byname []Person

func (x byname) Len() int           { return len(x) }
func (x byname) Less(i, j int) bool { return x[i].Name < x[j].Name }
func (x byname) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

func main() {
	people := []Person{
		{Name: "Rajesh", Age: 36},
		{Name: "Rakesh", Age: 32},
		{Name: "Ramesh", Age: 31},
		{Name: "Rabesh", Age: 30},
	}
	fmt.Println(people)
	sort.Sort(byage(people))
	fmt.Println(people)
	sort.Sort(byname(people))
	fmt.Println(people)

}
