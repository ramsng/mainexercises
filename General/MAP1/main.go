package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	f, err1 := os.Open("Map1.txt")
	if err1 != nil {
		log.Fatalf("Error reading file %v : %v", f, err1)
	}

	type Reader struct {
		contents string
	}

	_, err := ioutil.Readall
	if err != nil {
		log.Fatalf("Error reading file : %v", err)
	}

	//f = bufio.NewReader(rd io.)
	fmt.Println(contents)
	/*
		wordmap := make (map[string]int)
		for (err != nil) {

		}
	*/
}
