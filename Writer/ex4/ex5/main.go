package main

import (
	"bytes"
	"fmt"
)

func main() {
	bytebuff := []byte("Good day")
	bytebuff = append(bytebuff, []byte(" Gophers !")...)
	wrbuffer := bytes.NewBuffer(bytebuff)
	fmt.Println(wrbuffer)
	/*
		bytebuff = append(bytebuff, []byte("Good Day!"))
		wrbuffer.Write(bytebuff)
		file, err := os.Create("ex5.txt")
		if err != nil {
			log.Fatalf("Error open %s\t:%v", file, err)
		}
		file.Write(wrbuffer.Bytes())
	*/
}
