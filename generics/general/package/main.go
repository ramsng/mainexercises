package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type pack interface {
	constraints.Integer
}

func addfunc[t pack](a t, b t) (c t) {
	c = a + b
	return
}

func main() {
	fmt.Printf("Float 64 addition using constraints:%v", addfunc(80, 81))
}
