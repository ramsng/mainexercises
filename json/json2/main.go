package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var jsonblog = []byte(`[
		{"Name":"Mercedes Benz E-class","Model":2019,"Colors":["black","silver"]},
		{"Name":"Mercedes Benz C-class","Model":2021,"Colors":["grey","black","white"]}
		]`)
	type jstruct struct {
		Name   string
		Model  int
		Colors []string
	}
	var cars []jstruct
	err := json.Unmarshal(jsonblog, &cars)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(cars)
}
