package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the string:")
	data, err := reader.ReadBytes('\n')
	//data, err := reader.ReadBytes(' ')
	fmt.Println(err)
	data1 := string(data)
	fmt.Println(data1)
	//fmt.Println(data1)
	data2 := strings.Split(data1, " ")
	fmt.Println(data2[0])
	fmt.Println(data2[1])
	//fmt.Println(data2)
	//data, err = reader.ReadBytes(' ')
	//data, err := reader.ReadBytes(' ')
	/*
		fmt.Println(err)
		if err == io.EOF {
			fmt.Println(data)
			fmt.Println(string(data))
		}
			var finaldata []byte
			for {
				data, err = reader.ReadBytes(' ')
				fmt.Println("ERR1:", err)
				fmt.Println(data)
				if (err == io.EOF) || (err != nil) {
					break
				} else {
					finaldata = append(finaldata, data...)
				}
			}
			fmt.Println(finaldata)
			fmt.Println(string(finaldata))
	*/
}
