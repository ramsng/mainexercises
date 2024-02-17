package main

import "fmt"

// bi directional channels
func main() {
	// method 1 via channel processing
	ch1 := make(chan int)
	go func() {
		ch1 <- 42
		ch1 <- 44
		ch1 <- 43
	}()
	fmt.Println(<-ch1)
	fmt.Println(<-ch1)
	fmt.Println(<-ch1)
	// method 2 via channel processing
	ch2 := make(chan int, 2)
	ch2 <- 99
	ch2 <- 98
	fmt.Println(<-ch2)
	fmt.Println(<-ch2)
	// unidirectional channels
	//sendch := make(chan<- string)
	ch3 := make(chan int)
	go Sendonly(ch3)
	Recvonly(ch3)
	//recvch := make(<-chan <-chan string)
}

func Recvonly(recvch <-chan int) {
	fmt.Printf(" Channel String: %v ; Type: %T", <-recvch, recvch)
}

func Sendonly(sendch chan<- int) {
	sendch <- 24
}
