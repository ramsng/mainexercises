package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Average(a []int64) float64 {
	var cnt int = 0
	var tot, tot1 int64
	for cnt, tot1 = range a {
		tot = tot + tot1
		//	fmt.Println(tot1)
	}
	//fmt.Println("count, Total :", tot, cnt)
	avg := tot / (int64(cnt) + 1)
	return float64(avg)
}

func main() {
	var input []int64
	fmt.Println("Enter the inputs : ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	//fmt.Println(scanner.Text())
	str := strings.Split(scanner.Text(), " ")
	for _, b := range str {
		//	fmt.Println(a, b)
		c, _ := strconv.ParseInt(b, 10, 64)
		//	fmt.Println("strconv values", c)
		input = append(input, c)
	}
	fmt.Println("Average on the inputs : ", input)
	fmt.Println("is :", Average(input))
}
