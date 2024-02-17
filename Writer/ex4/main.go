package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func main() {
	file, err := os.Create("ex4.txt")
	if err != nil {
		log.Fatalf("Error opening file %v\t:%v", file, err)
	}
	type data struct {
		statement string
		writets   string
	}

	var data1 data = data{
		statement: "Hello to Gophers!",
		writets:   time.Now().String(),
	}

	var writedata string = data1.statement + data1.writets
	writebyte := []byte("Hello Google")
	fmt.Println(writedata)
	fmt.Println(string(writebyte))
	_, err = io.WriteString(file, string(writedata))
	if err != nil {
		log.Fatalf("Error writing file %v\t:%v", file, err)
	}
}
