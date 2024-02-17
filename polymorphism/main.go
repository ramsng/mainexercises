package main

import "fmt"

type person struct {
	Name  string
	actor bool
}

type act interface {
	bond()
}

func (p person) bond() {
	fmt.Printf("I am %v, Actor experience: %v\n", p.Name, p.actor)
}
func experience(a act) {
	a.bond()
}
func main() {
	p1 := person{
		Name:  "James Bond",
		actor: true,
	}
	p2 := person{
		Name:  "Berk",
		actor: false,
	}

	experience(p1)
	experience(p2)
}
