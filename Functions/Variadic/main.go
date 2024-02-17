package main

import (
	"fmt"
	"strings"
)

func main() {
	//	mapb := []string{}
	mapc := icecream()
	mapd := delicious("DELICIOUS", mapc...)

	for a, b := range mapd {
		defer fmt.Printf("%v\n", b)
		fmt.Printf("DELICIOUS ICECREAMS: %d\n", a+1)
	}
}

func icecream() []string {
	mapa := []string{"Vanilla", "Chocomonco", "Mango Safron", "blueberry butterscotch"}
	return mapa
}
func delicious(flag string, mape ...string) []string {
	mapa := []string{}
	if flag == "DELICIOUS" {
		for _, b := range mape {
			if (b == "Vanilla") || (b == "Mango Safron") {
				mapa = append(mapa, strings.ToUpper(b))
				//				fmt.Println(b, flag)
			}
		}
	}
	return mapa
}
