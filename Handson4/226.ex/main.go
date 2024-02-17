package main

import (
	"fmt"
	"sort"
)

func EvenOdd(v int, Oddchan, Evenchan chan int) {

	//evenarr, oddarr := []int
	if v%2 > 0 {
		Oddchan <- v
	} else {
		Evenchan <- v
	}
	//return Evenchan, Oddchan
}

func Newchan(num1, num2 int) chan int {
	Newchan := make(chan int)
	go func() {
		for i := num1; i < num2; i++ {
			Newchan <- i
			//fmt.Println("gen values", i)
		}
		close(Newchan)
	}()
	return Newchan
}

func Fanin(Oldchan, Newchan, Chan1, Chan2 chan int) {
	var break1sw, break2sw string
	//	defer ccall()
	//defer close(Oddchan)
	func() {
		for {
			select {
			case v, ok := <-Oldchan:
				if !ok {
					break1sw = "Y"
				} else {
					EvenOdd(v, Chan1, Chan2)
				}
			case v, ok := <-Newchan:
				if !ok {
					break2sw = "Y"
				} else {
					EvenOdd(v, Chan1, Chan2)
				}
			}
			if break1sw == "Y" && break2sw == "Y" {
				close(Chan1)
				close(Chan2)
				break
			} else {
				continue
			}
		}
	}()
}

func main() {
	Evenchan := make(chan int)
	Oddchan := make(chan int)
	var evenarr, oddarr []int
	var break1, break2 string
	go Fanin(Newchan(0, 10), Newchan(11, 21), Oddchan, Evenchan)
	for {
		select {
		case v, ok := <-Evenchan:
			if !ok {
				break1 = "Y"
			} else {
				evenarr = append(evenarr, v)
				//	fmt.Println("Even #", v)
			}
		case v, ok := <-Oddchan:
			if !ok {
				break2 = "Y"
			} else {
				oddarr = append(oddarr, v)
			}
		}
		if break1 == "Y" && break2 == "Y" {
			sort.Ints(oddarr)
			sort.Ints(evenarr)
			fmt.Println("Odd Numbers :", oddarr)
			fmt.Println("Even Numbers :", evenarr)
			break
		}
	}
}
