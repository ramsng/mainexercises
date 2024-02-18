package main

import (
	"fmt"

	"github.com/ramsng/sub01/tree/main/dog"
)

func main() {
	fmt.Println("Enter your age : ")
	var a int
	if _, err := fmt.Scan(a); err != nil {
		fmt.Errorf("Scan errors", err)
	}
	a = dog.Years(a)
	fmt.Println("Your Dog's age : ")
}
