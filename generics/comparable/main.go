package main

import "fmt"

func compareadd[k comparable, v int | float64](m map[k]v) v {
	var s v = 0
	for _, v := range m {
		s += v
	}
	return s
}
func main() {
	m := map[string]int{
		"first":  35,
		"second": 35,
		"third":  30,
	}

	m1 := map[string]float64{
		"first":  35.25,
		"second": 35.38,
		"third":  30.45,
	}
	fmt.Println("Adding integers : ", compareadd(m))
	fmt.Println("Adding floats : ", compareadd(m1))
}
