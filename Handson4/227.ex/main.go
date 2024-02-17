package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg1 sync.WaitGroup

func main() {
	ch1 := make(chan int)
	go send(ch1)
	wg1.Add(1)
	go receive(ch1)
	/*go func() {
		for v := range ch1 {
			fmt.Println("Value :", v)
		}
		wg1.Done()
	}()
	*/
	wg1.Wait()
	//wg1.Done()
}

func send(ch1 chan int) {
	//go func() {
	for i := 1; i < 10; i++ {
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		value := i * 1000
		fmt.Println("Send: ", value)
		ch1 <- value
	}
	//	ch1 <- 999
	close(ch1)
	//}()
	//return ch1
}

func receive(ch1 <-chan int) {
	//var wg sync.WaitGroup
	//wg.Add(1)
	for v := range ch1 {
		fmt.Println("Value :", v)
		//	wg.Wait()
	}
	wg1.Done()
}
