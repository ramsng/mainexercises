package main

import (
	"fmt"
	"sort"
	"strings"
)

var a = []string{"Ranjish", "Ramesh", "Suresh", "Rajesh", "Rocklyn", "Raahul"}
var findstring = "Roger"

func compare(n int) int {
	i, found := sort.Find(n, func(i int) int {
		return strings.Compare(findstring, a[i])
	})
	if found {
		fmt.Println("String found at entry :", i+1)
	} else {
		fmt.Println("index:", i)
		y := []string{}
		y = append(y, a[(i):]...)
		fmt.Println("y:", y)
		a = append(a[:i], findstring)
		fmt.Println("append string:", a)
		a = append(a, y[:]...)

		//a = append(a[:i], findstring, a[i+1])
		fmt.Println("Final string:", a)
		fmt.Println("String inserted at entry :", i+1)
	}
	return i
}

func main() {
	fmt.Println(a)
	fmt.Println("-- Post SORT ---")
	sort.Strings(a)
	fmt.Println(a)
	x := compare(6)
	fmt.Println("-- Returned value ---")
	fmt.Println(x)
	fmt.Println(a)
}
