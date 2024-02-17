package main

import "fmt"

// 4*3*2*1
//var b int = 1

func factorial(a int) int {
	if a == 1 {
		return a
	}
	return a * factorial(a-1)
}

func main() {
	fmt.Println("Factorial for the value 10 using functions : ", factorial(10))
	var factval int = 1
	for fact := 10; fact >= 1; fact-- {
		factval *= fact
		//	factval = fact * factval
	}
	fmt.Println("Factorial for the value 10 using Loops     : ", factval)
}
