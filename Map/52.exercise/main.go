package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	xa := []string{`Short`, `stories`, `have`, `no`, `set`, `length`, `In`, `terms`, `of`, `word`, `count`, `there`, `is`, `no`, `official`, `demarcation`, `between`, `an`, `anecdote`, `a`, `short`, `story`, `and`, `a`, `novel`, `Like`, `the`, `novel,`, `the`, `short`, `story`, `predominant`, `shape`, `reflects`, `the`, `demands`, `of`, `the`, `available`, `markets`, `for`, `publication,`, `and`, `the`, `evolution`, `of`, `the`, `form`, `seems`, `closely`, `tied`, `to`, `the`, `evolution`, `of`, `the`, `publishing`, `industry`, `and`, `the`, `submission`, `guidelines`, `of`, `its`, `constituent`, `houses`}
	mapa := map[string]int{}
	for _, b := range xa {
		mapa[strings.ToLower(b)]++
	}
	sort.Ints(mapa)
	for a, b := range mapa {
		fmt.Printf("%v\t\t\t%v\n", a, b)
	}
}
