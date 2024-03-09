// Package Average is the customized package on Average
package avg

import "sort"

//Function Centeredavg removes the first and last integer in the sorted array and calcualtes the average
func Centeredavg(xi []int) float64 {
	sort.Ints(xi)
	xi = xi[1:(len(xi) - 1)]
	var a int
	for _, v := range xi {
		a += v
	}
	return float64(a) / float64(len(xi))
}
