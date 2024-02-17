package main

import "fmt"

type count string

func main() {
	fmt.Println("** Practicing anounymous func **")
	x := func() string {
		return fmt.Sprintln("Return from anonymous func")
	}()
	fmt.Printf("\nVerify in main module :\t %v", x)

}
