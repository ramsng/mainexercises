package main

// Channel range //
import "fmt"

func main() {
	ch1 := make(chan int)
	go sendonly(ch1)
	for a := range ch1 {
		fmt.Println(a)
	}
}
func sendonly(ch1 chan<- int) {
	for i := 0; i < 10; i++ {
		ch1 <- i
	}
	close(ch1)
}
