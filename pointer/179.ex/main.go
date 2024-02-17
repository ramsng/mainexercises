package main

import "fmt"

/*
Hands-on exercise #75 - dereference an address
In the provided code, there are variables that store VALUES of TYPE *int and TYPE *string. The
values stored are memory addresses. Using the variables provided, do the following:
● print the VALUE stored in each variable
○ these will be memory addresses
● print the TYPE of each variable
● print the data stored at memory locations
○ dereference the stored memory address *
*/

var (
	a, b, c *string
	d       *int
)

func init() {
	p := "Drop by drop, the bucket gets filled."
	q := "Persistently, patiently, you are bound to succeed."
	r := "The meaning of life is ..."
	n := 42
	a = &p
	b = &q
	c = &r
	d = &n
	fmt.Println("Values of variables  p : ", p)
	fmt.Println("Values of variables  q : ", q)
	fmt.Println("Values of variables  r : ", r)
	fmt.Println("Values of variables  n : ", n)
}

func main() {
	// Print the VALUE stored in each variable
	fmt.Println("Values in  a     : ", a)
	fmt.Println("Values in  b     : ", b)
	fmt.Println("Values in  c     : ", c)
	fmt.Println("Values in  d     : ", d)
	// print the TYPE of each variable
	fmt.Printf("Value types in  a : %T\n", a)
	fmt.Printf("Value types in  b : %T\n", b)
	fmt.Printf("Value types in  c : %T\n", c)
	fmt.Printf("Value types in  d : %T\n", d)
	// print the data stored at memory locations
	fmt.Println("Values of variables  a : ", *a)
	fmt.Println("Values of variables  b : ", *b)
	fmt.Println("Values of variables  c : ", *c)
	fmt.Println("Values of variables  d : ", *d)
}
