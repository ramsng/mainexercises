package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func subtract(a int, b int) int {
	return a - b
}

func compute(a int, b int, c func(int, int) int) int {
	return c(a, b)
}

func main() {
	a := compute(15, 10, add)
	fmt.Println(a)
	a = compute(15, 10, subtract)
	fmt.Println(a)
}
