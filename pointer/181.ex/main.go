package main

import "fmt"

type field struct {
	name string
}

func valsem(f field) field {
	f.name = "Henry"
	return f
}
func pointersem(f *field) {
	f.name = "Louis"
}
func main() {
	var name field = field{"Renold"}
	fmt.Printf("\nInitial Value        :%v\ttype:%T", name.name, name.name)
	fmt.Printf("\nPost Value semantic  :%v\ttype:%T", name.name, name.name)
	name = valsem(name)
	fmt.Printf("\nPost Value semantic return :%v\ttype:%T", name.name, name.name)
	pointersem(&name)
	fmt.Printf("\nPost Pointer semantic:%v\ttype:%T", name.name, name.name)
}
