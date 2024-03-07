package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ramsng/sub01/dog"
)

func input(p1 *int) {
	fmt.Fprint(os.Stdout, "Enter your age : ")
	_, err := fmt.Scan(p1)
	if err != nil {
		log.Println(err)
		*p1 = 0
	}
}

func main() {
	var p1 int = 0
	input(&p1)
	fmt.Println("Entered value is ", p1)
	a := dog.Years(&p1)
	fmt.Println("Your Dog's age : ", a)
}

/*
func main() {
	input := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter the value here to display :")
	input.Scan()
	//	if input.Text() ==
	fmt.Println("Entered value is ", input.Text())
	fmt.Println("Entered value is ", string(input.Bytes()))
	//fmt.Println("Entered value is ", input.Buffer())
}
*/
//reader := bufio.NewReader(os.Stdin)

/*
	line, err := reader.ReadBytes('1')
	if err != nil {
		err1 := errors.New("New error")
		fmt.Println(err1)
	}
	b := string(line)
	fmt.Println("Entered value is ", b)
	a, err := strconv.Atoi(b)
*/
//var line *int

/*if err != nil {
	log.Println(err)
	//fmt.Println(err1)
}
*/
