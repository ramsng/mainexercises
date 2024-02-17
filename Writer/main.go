package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println(" Writer interface pgm")
	x, err := os.Create("output.txt")
	defer x.Close()
	if err != nil {
		log.Fatal("error at open", err)
	}
	s := []byte("Hello Gophers!!")
	_, err = x.Write(s)
	if err != nil {
		log.Fatal("error at write", err)
	}
}
