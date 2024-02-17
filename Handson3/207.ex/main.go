package main

import (
	"fmt"
	"runtime"
	"sync"
)

func mirror1(a float64) {
	runtime.Gosched()
	fmt.Println("counter from mirror1 :\t", a)
	fmt.Println("Number of GO calls: ", runtime.NumGoroutine())
	wg.Done()
}

func mirror2(a float64) {
	fmt.Println("counter from mirror2 :\t", a)
	wg.Done()
}

var wg sync.WaitGroup

func main() {
	fmt.Println("Number of CPU's: ", runtime.NumCPU())
	fmt.Println("Number of GO calls: ", runtime.NumGoroutine())
	for i := 0; i < 25; i++ {
		wg.Add(1)
		go mirror1(float64(i))
		/*    wg.Add(i)
		go mirror2(float64(i))
		fmt.Println("Number of CPU's: ", runtime.NumGoroutine())
		*/
	}
	wg.Wait()
	fmt.Println("Number of GO calls Finally: ", runtime.NumGoroutine())
}
