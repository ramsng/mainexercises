package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type User struct {
	Name string
	Age  int
	Auth bool
}

func main() {
	u1 := User{
		Name: "Rams",
		Age:  36,
		Auth: true,
	}
	u2 := User{
		Name: "Ravi",
		Age:  32,
		Auth: false,
	}
	multiuser := []User{u1, u2}
	u1json, err := json.Marshal(u1)
	if err != nil {
		log.Fatal("error in json", err)
	}
	u2json, err := json.Marshal(u2)
	if err != nil {
		log.Fatal("error in json", err)
	}
	mujson, err := json.Marshal(multiuser)
	if err != nil {
		log.Fatal("error in json", err)
	}
	fmt.Println(u1)
	fmt.Println(string(u1json))
	fmt.Println(u2)
	fmt.Println(string(u2json))
	fmt.Println(multiuser)
	fmt.Println(string(mujson))
}
