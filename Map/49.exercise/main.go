package main

import "fmt"

func main() {
	/*
		mapi := map[string]string{
			"bond_james": "`shaken,not stirred`, `martinis`, `fast cars`",
		}
		var maxsize int
		mapi["moneypenny_jenny"] = "`intelligence`,`literature`,`Computer Science`"
		mapi["no_dr"] = "`cats`,`ice cream`,`sunsets`"
		for a, _ := range mapi {
			if maxsize < len(a) {
				maxsize = len(a)
			}
		}
		//fmt.Println("Maximum size :", maxsize)
		for i, j := range mapi {
			//fmt.Println("size before processing: \n", len(i))
			if size := maxsize - len(i); size > 0 {
				//	fmt.Println("size after compute: \n", size, maxsize)
				for k := 0; k < size; k++ {
					i = i + " "
				}
				//	fmt.Println("size after pocessing :\n", len(i))
			}
			fmt.Printf(" Person : %v\tLikes : %v\n", i, j)
		}
	*/
	/*
		xa := make([]string, 3, 3)
		xb := make([]string, 3, 3)
		xc := make([]string, 3, 3)
		xa[0] = "`shaken,not stirred`"
		xa[1] = "`martinis`"
		xa[2] = "`fast cars`"
	*/
	mapi := map[string][]string{}
	mapi[`bond_James`] = []string{"`shaken,not stirred`", "`martinis`", "`fast cars`"}
	mapi[`moneypenny_jenny`] = []string{"`intelligence`", "`literature`", "`Computer Science`"}
	mapi[`no_dr`] = []string{"`cats`", "`ice cream`", "`sunsets`"}
	mapi[`fleming_ian`] = []string{"`steaks`", "`cigars`", "`espionage`"}
	/*	xb = append(xb, "`intelligence`", "`literature`", "`Computer Science`")
		mapi[`moneypenny_jenny`] = xb
		xc = append(xc, "`cats`", "`ice cream`", "`sunsets`")
		mapi[`no_dr`] = xc
	*/
	for a, b := range mapi {
		fmt.Printf("Person : %v - Their likes \n", a)
		for c, d := range b {
			fmt.Printf("%d . %v\n", c+1, d)
		}
	}
	delete(mapi, `no_dr`)
	fmt.Println("** Updated List **")
	for a, b := range mapi {
		fmt.Printf("Person : %v - Their likes \n", a)
		for c, d := range b {
			fmt.Printf("%d . %v\n", c+1, d)
		}
	}
}
