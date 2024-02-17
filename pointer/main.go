package main

import "fmt"

/* Hands-on exercise #74 - take an address to create a pointer
● Create a value and assign it to a variable.
● Print the address of that value.
https://go.dev/play/p/rvcrIhHRntg
curriculum item # 176-hands-on-74

func main() {
	a := 42
	fmt.Println("Value   : ", a)
	fmt.Println("Address : ", &a)
}
*/
var (
	a, b, c *int
	e, f    *string
	g, h    int
)

func main() {
	fmt.Println("Address : ", &a)
	fmt.Println("Address : ", &b)
	fmt.Println("Address : ", &c)
	fmt.Println("Address : ", &e)
	fmt.Println("Address : ", &f)
	fmt.Println("Address : ", &g)
	fmt.Println("Address : ", &h)
	fmt.Println("")
	fmt.Println("Value : ", a)
	fmt.Println("Value : ", b)
	fmt.Println("Value : ", c)
	fmt.Println("Value : ", e)
	fmt.Println("Value : ", f)
	fmt.Println("Value : ", g)
	fmt.Println("Value : ", h)
	fmt.Println("")
}
