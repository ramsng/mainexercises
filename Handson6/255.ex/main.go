package main

import (
	"fmt"

	"github.com/ramsng/mainexercises/Handson6/255.ex/avg"
	"github.com/ramsng/mainexercises/Handson6/255.ex/input"
)

func main() {
	fmt.Println(input.Geninput())
	for _, v := range input.Geninput() {
		fmt.Println(avg.Centeredavg(v))
	}
}
