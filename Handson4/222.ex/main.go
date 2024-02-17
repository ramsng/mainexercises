package main

import "fmt"

func main() {
	var arr = []string{"Hello", "World", "Gophers", "exercises"}
	var ch1 = make(chan string)
	go func(arr []string) {
		for _, v := range arr {
			ch1 <- v
		}
		close(ch1)
	}(arr)
	fmt.Println("Array to Channel strings :")
	/*for v := range ch1 {
		fmt.Println(v)
	}
	*/
	for {
		v, ok := <-ch1
		if ok {
			fmt.Println(v)
		} else {
			break
		}
	}
}
