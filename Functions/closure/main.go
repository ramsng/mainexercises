package main

import "fmt"

func increamentor() func() int {
	var a int = 0
	return func() int {
		a++
		return a
	}
}
func main() {
	fmt.Println(increamentor()())
	inc := increamentor()
	fmt.Println(inc())
	inc = increamentor()
	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(inc())
}
