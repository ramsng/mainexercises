package main

import "fmt"

type person struct {
	first       string
	last        string
	icecreamcnt int
	donutcnt    int
}

func (p person) deserts() {
	fmt.Printf(" %v %v says 'Thanks for %d icecreams & %v donuts\n", p.first, p.last, p.icecreamcnt, p.donutcnt)
}
func main() {
	p1 := person{
		first:       "Ezhil",
		last:        "Rams",
		icecreamcnt: 2,
		donutcnt:    2,
	}
	p2 := person{
		first:       "Vedu",
		last:        "Rams",
		icecreamcnt: 2,
		donutcnt:    2,
	}
	p3 := person{
		first:       "Gayu",
		last:        "Murgs",
		icecreamcnt: 2,
		donutcnt:    2,
	}
	p1.deserts()
	p2.deserts()
	p3.deserts()

}
func init() {
	fmt.Println("*** Hi to Rams family! ***")
}
