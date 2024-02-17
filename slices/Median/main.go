package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("** Iteration 1 **")
	xi := []float64{42, 45, 43, 61, 69}
	fmt.Println(findmedian(xi))
	fmt.Println(xi)
	fmt.Println("** Iteration 2 **")
	xj := []float64{49, 44, 46, 48}
	fmt.Println(findmedian(xj))
	fmt.Println(xj)
	fmt.Println("** Iteration 3 **")
	xk := []float64{43, 45, 49, 48}
	fmt.Println(findmedian(xk))
	fmt.Println(xk)
	fmt.Println("** Iteration 4 **")
	xl := []float64{42, 45, 43, 61, 69}
	fmt.Println(findmedian2(xl))
	fmt.Println(xl)
	fmt.Println("** Iteration 5 **")
	xm := []float64{49, 44, 46, 48}
	fmt.Println(findmedian2(xm))
	fmt.Println(xm)
	fmt.Println("** Iteration 6 **")
	xn := []float64{43, 45, 49, 48}
	fmt.Println(findmedian2(xn))
	fmt.Println(xn)
}

func findmedian(x []float64) float64 {
	//	fmt.Println(x)
	sort.Float64s(x)
	//	fmt.Println(x)
	y := len(x) % 2
	//fmt.Println(y)
	if y == 1 {
		return x[(len(x) / 2)]
	} else {
		z := x[(len(x)/2)-1] + x[(len(x)/2)]
		return z / 2
	}
}
func findmedian2(x []float64) float64 {
	//	fmt.Println(x)
	w := make([]float64, 4, 4)
	copy(w, x)
	sort.Float64s(w)
	//	fmt.Println(x)
	y := len(w) % 2
	//fmt.Println(y)
	if y == 1 {
		return w[(len(w) / 2)]
	} else {
		z := w[(len(w)/2)-1] + w[(len(x)/2)]
		return z / 2
	}
}
