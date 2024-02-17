package main

import "fmt"

type person struct {
	name     string
	icecream string
	count    int
}

func (p person) String() string {
	return fmt.Sprintf("I am %v, ate %v %v icecreams\n", p.name, p.count, p.icecream)
}

func appender(s fmt.Stringer) {
	fmt.Println("Hello Guys! " + s.String())
}

func main() {
	p1 := person{
		name:     "Raju",
		icecream: "Vanilla",
		count:    5,
	}
	p2 := person{
		name:     "Ravi",
		icecream: "Chocolate",
		count:    2,
	}
	fmt.Println(p1)
	fmt.Println(p2)
	appender(p1)

}
