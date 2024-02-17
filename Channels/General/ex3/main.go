package main

import "fmt"

func main() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan int)
	go evenodd(even, odd, quit)
	evenl, oddl := selectcase(even, odd, quit)
	fmt.Println("-----------------")
	fmt.Println("Even numbers < 10")
	for _, c := range evenl {
		fmt.Println(c)
	}
	fmt.Println("-----------------")
	fmt.Println("Odd numbers < 10")
	for _, c := range oddl {
		fmt.Println(c)
	}
}

func evenodd(e, o, q chan<- int) {
	for i := 0; i < 10; i++ {
		switch {
		case i%2 == 0:
			{
				e <- i
				//			fmt.Println("even:", i)
			}
		case i%2 == 1:
			{
				o <- i
				//			fmt.Println("odd:", i)
			}
		}
	}
	//close(e)
	//close(o)
	q <- 0
	close(q)
}

func selectcase(e, o, q <-chan int) (a, b []int) {
	for {
		select {
		case v := <-e:
			a = append(a, v)
		case v := <-o:
			b = append(b, v)
		case v, ok := <-q:
			if !ok {
				fmt.Println(" close case not OK ", v, ok)
				return a, b
			} else {
				fmt.Println(" close case OK ", v, ok)
			}
		}
	}
}
