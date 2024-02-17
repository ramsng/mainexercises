package main

import "fmt"

type person struct {
	first  string
	last   string
	rollno int
	rank   int
}

func main() {
	p1 := person{
		first:  "Arvind",
		last:   "Aropkiaraj",
		rollno: 60056,
		rank:   6,
	}
	p2 := person{
		first:  "Ezhil",
		last:   "Arasan",
		rollno: 60040,
		rank:   3,
	}
	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p1.first, p2.last)

}
