package main

import "fmt"

/*
func addmodel[t int | float64](a, b t) t {
	return a + b
}
*/
type generic interface {
	~int | ~float64 | ~complex64
}

func addmodel[t generic](a, b t) t {
	return a + b
}

// type Alias

type mytype float64

func main() {
	// add float values
	fmt.Printf("Value for add: %v\n", addmodel(42.2, 84.56))
	// add int values
	fmt.Printf("Value for add: %v\n", addmodel(42, 84))
	// add float values => using interface
	fmt.Printf("Value for add: %v\n", addmodel(42.2, 84.56))
	// add int values => using interfaces
	fmt.Printf("Value for add: %v\n", addmodel(42, 84))
	const var1 mytype = 40.14
	fmt.Printf("Add model for alias: %v\n", addmodel(var1, 84.56))
}
