package main

import (
	"bytes"
	"fmt"
)

func main() {
	//	b := bytes.NewBufferString("Hello Gophers!!!")
	b := bytes.NewBufferString(" ")
	fmt.Println(b.String())
	b.WriteString("Hello Gophers!!! \nEnjoy the Writers interface section ")
	fmt.Println(b.String())
	b.Reset()
	fmt.Println(b.String())
	b.WriteString("Enjoy the Writers interface section - post reset\n")
	//fmt.Println(b.String())
	fmt.Println(fmt.Stringer(b))
}
