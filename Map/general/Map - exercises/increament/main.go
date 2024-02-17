package main

import "fmt"

func main() {
	mapi := map[string]int{}
	fmt.Println("* * initial Map * *")
	fmt.Println(mapi)
	fmt.Println(len(mapi))
	fmt.Println("* * Revised Map * *")
	mapi["beautiful"] = 40
	fmt.Println(mapi)
	fmt.Println(len(mapi))
	fmt.Println("* * Final map * *")
	mapi["beautiful"]++
	mapi["wonderful"]++
	fmt.Println(mapi)
	fmt.Println(len(mapi))
}
