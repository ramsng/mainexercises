package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func readfile(f string) []byte {
	//	var read []byte
	read, err := os.ReadFile(f)
	if err != nil {
		log.Fatalf("Read error : %v", err)
	}
	return read
}
func readtime() {
	start := time.Now()
	fmt.Println(string(readfile("exercise1.txt")))
	fmt.Println(time.Since(start))
	time.Sleep(2 * time.Second)
}
func main() {
	start := time.Now()
	readtime()
	fmt.Println("\n", time.Since(start))
}
