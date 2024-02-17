package main

import "fmt"

func main() {
	// exercise 42 - create array and display all the vaiable
	/*
		fmt.Println("Exercise 42 - ARRAY CREATE AND DISPLAY ALL THE VALUES")
		xa := [5]float64{}
		for i := 0; i < 5; i++ {
			j := float64(i)
			xa[i] = math.Pow(j, j)
		}
		for i, j := range xa {
			fmt.Println(i+1, j)
		}
	*/
	/*	Exercise43 - create slice and display*/
	// create slice
	//xi := make([]int, 0, 5)
	//xi = append(xi, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52)
	//sort.Ints(xi)
	/*
		for a, b := range xi {
			fmt.Printf("Index : %#v\tValue : %#v\n", a, b)
		}
			Exercise44 - slicing the slice*/
	/*
		fmt.Printf("%#T\t%v\n", xi, xi)
		xa := make([]int, 0, 10)
		xa = append(xa, xi[:5]...)
		fmt.Printf("%#T\t%v\n", xa, xa)
		xb := make([]int, 0, 10)
		xb = append(xb, xi[5:5+5]...)
		fmt.Printf("%#T\t%v\n", xb, xb)
		xc := make([]int, 0, 10)
		xc = append(xc, xi[2:2+5]...)
		fmt.Printf("%#T\t%v\n", xc, xc)
		xd := make([]int, 0, 10)
		xd = append(xd, xi[1:1+5]...)
		fmt.Printf("%#T\t%v\n", xd, xd)
	*/
	// Exercise: 45
	/*
		x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52}
		x = append(x, 52)
		fmt.Printf("%T\tLength : %v\tValue : %v\n", x, len(x), x)
		x = append(x, 53, 54, 55)
		fmt.Printf("%T\tLength : %v\tValue : %v\n", x, len(x), x)
		y := make([]int, 0)
		y = append(x, 56, 57, 58, 59, 60)
		fmt.Printf("%T\tLength : %v\tValue : %v\n", y, len(y), y)
	*/
	// Exercise: 46
	/*
		x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
		fmt.Println("Initial slice : ", x)
		x = append(x[:3], x[6:]...)
		fmt.Println("updated slice : ", x)
	*/
	//Exercise:47
	/*
		x := make([]string, 0)
		x = append(x, ` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`)
		fmt.Printf("Total States in USA : %d\n", len(x))
		for a, b := range x {
			fmt.Println(a+1, b)
		}
		fmt.Printf("\nCapacity of array : %d", cap(x))
	*/
	//Exercise:48
	ai := []string{"James", "Bond", "Shaken, not stirred"}
	aj := []string{"Miss", "Moneypenny", "I'm 008."}
	ak := [][]string{ai, aj}
	fmt.Printf("\n%v", ak[0])
	fmt.Printf("\n%v", ak[1])
}
