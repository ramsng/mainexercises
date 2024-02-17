package main

import (
	"fmt"
	"sort"
)

func main() {

	var a = []int{40, 20, 30, 24, 53, 27, 30}
	for _, v := range a {
		fmt.Println(v)
	}
	fmt.Println("--- Post sort ---")
	sort.Ints(a)
	for _, v := range a {
		fmt.Println(v)
	}
}
