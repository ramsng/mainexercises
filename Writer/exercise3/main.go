package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"strconv"
)

type count int

func (c count) writeio(w io.Writer) {
	w.Write([]byte(c.String()))
}

func (c count) String() string {
	x := fmt.Sprint("\nMy lucky number is :", (strconv.Itoa(int(c))))
	return x
}

func main() {
	f, err := os.Create("exercise3.txt")
	if err != nil {
		log.Fatalf("Error opening the file %v: %v", f, err)
	}
	defer f.Close()
	b := bytes.NewBuffer([]byte("** Lucky number draw ***"))
	f.Write([]byte("** Lucky number draw ***"))
	if err != nil {
		log.Fatalf("Error initial write of file %v: %v", f, err)
	}
	var nbr count = count(rand.Intn(100))
	nbr.writeio(f)
	if err != nil {
		log.Fatalf("Error writing the file %v: %v", f, err)
	}
	nbr.writeio(b)
	if err != nil {
		log.Fatalf("Error writing the file %v: %v", b, err)
	}
	fmt.Println(b)
}
