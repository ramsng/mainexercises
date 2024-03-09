package main

import (
	"fmt"

	"github.com/ramsng/mainexercises/Handson6/254.ex/quote"
	"github.com/ramsng/mainexercises/Handson6/254.ex/word"
)

func main() {
	fmt.Println(quote.SunAlso)
	fmt.Println(word.Count(quote.SunAlso))
	fmt.Println(quote.SunAlso)
	for m, v := range word.Wordcount(quote.SunAlso) {
		fmt.Println(m, v)
	}
	for m, v := range word.Wrdcount(quote.SunAlso) {
		fmt.Println(m, v)
	}
}
