package main

import (
	"encoding/json"
	"fmt"
	"os"
)

//JSON basics marshalling & unmarshalling

type Jstruct struct {
	Name   string
	Model  int
	Colors []string
}

func main() {
	car1 := Jstruct{
		Name:   "Mercedes Benz E-class",
		Model:  2019,
		Colors: []string{"black", "silver"},
	}
	car2 := Jstruct{
		Name:   "Mercedes Benz C-class",
		Model:  2021,
		Colors: []string{"grey", "black", "white"},
	}
	b, err := json.Marshal(car1)
	if err != nil {
		fmt.Println("JSON Marshall error", err)
	}
	os.Stdout.Write(b)
	b, err = json.Marshal(car2)
	if err != nil {
		fmt.Println("JSON Marshall error", err)
	}
	os.Stdout.Write(b)
}
