package main

import (
	"log"
	"os"
)

func main() {
	f, err := os.Create("exercise2.txt")
	if err != nil {
		log.Fatalf("Error opening file %v: %v", f, err)
	}
	defer f.Close()
	f.WriteString("Hello World!!")
	f.WriteString("\nEnjoy the new GO sessions!")
	b := []byte("\nGo through the Introductory part")
	f.Write(b)
}
