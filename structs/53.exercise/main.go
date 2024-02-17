package main

import "fmt"

func main() {
	type person struct {
		first    string
		last     string
		icecream []string
	}
	p1 := person{
		first:    "Rams",
		last:     "Azhags",
		icecream: []string{"Strawberry", "Vanilla", "Chocochips"},
	}
	p2 := person{
		first:    "Gayu",
		last:     "Murugan",
		icecream: []string{"Chocolate", "Mango"},
	}
	fmt.Printf("Name: %v %v\t & their favoirite icecreams\n", p1.first, p1.last)
	for a, b := range p1.icecream {
		fmt.Printf("%v . %v\n", a+1, b)
	}
	fmt.Printf("Name: %v %v\t & their favoirite icecreams\n", p2.first, p2.last)
	for a, b := range p2.icecream {
		fmt.Printf("%v . %v\n", a+1, b)
	}
}
