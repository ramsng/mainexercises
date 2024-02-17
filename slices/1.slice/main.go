package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	/*
		xi := []string{"Hello", "Gophers", "World"}
		fmt.Printf("\n%v\n%d\n", xi, len(xi))
		for a, b := range xi {
			fmt.Printf("%v\t%v\n", a+1, b)
		}
		c,d := range(xi)
		for i=0,i<c;i++ {
			fmt.Printf("%v\t%v\n", i+1, d)
		}
	*/
	xj := []string{
		"Almond Biscotti Café",
		"Banana Pudding ",
		"Balsamic Strawberry (GF)",
		"Bittersweet Chocolate (GF)",
		"Blueberry Pancake (GF)",
		"Bourbon Turtle (GF)",
		"Browned Butter Cookie Dough",
		"Chocolate Covered Black Cherry (GF)",
		"Chocolate Fudge Brownie",
		"Chocolate Peanut Butter (GF)",
		"Coffee (GF)",
		"Cookies & Cream",
		"Fresh Basil (GF)",
		"Garden Mint Chip (GF)",
		"Lavender Lemon Honey (GF)",
		"Lemon Bar",
		"Madagascar Vanilla (GF)",
		"Matcha (GF)",
		"Midnight Chocolate Sorbet (GF, V)",
		"Non-Dairy Chocolate Peanut Butter (GF, V)",
		"Non-Dairy Coconut Matcha (GF, V)",
		"Non-Dairy Sunbutter Cinnamon (GF, V)",
		"Orange Cream (GF) ",
		"Peanut Butter Cookie Dough",
		"Raspberry Sorbet (GF, V)",
		"Salty Caramel (GF)",
		"Slate Slate Different",
		"Strawberry Lemonade Sorbet (GF, V)",
		"Vanilla Caramel Blondie",
		"Vietnamese Cinnamon (GF)",
		"Wolverine Tracks (GF)"}
	fmt.Printf(" Ice cream varieties - %d", len(xj))
	for c, d := range xj {
		fmt.Printf("\n %d.\t%#v", c, d)
	}
	for i := 1; i <= 5; i++ {
		fmt.Printf("\nMy Lucky icecream on round %d is ------ %d.\t%#v;", i, rand.Intn(len(xj))+1, xj[i])
	}
	fmt.Println("\n*************************************")
	fmt.Println("\nAdding few more icecreams to the menu")
	fmt.Println("\n*************************************")
	xj = append(xj, "Crunchy Choco", "Vanilla rose", "Safron Vanilla")
	fmt.Printf("\nHere is the updated list on ice cream varieties %d", len(xj))
	for i := 0; i < len(xj); i++ {
		fmt.Printf("\n%d\t%#v", i+1, xj[i])
	}
	fmt.Println("\n*************************************")
	fmt.Println("\n Reducing icecreams from the menu")
	fmt.Println("\n*************************************")
	xj = append(xj[0:4], xj[16:]...)
	fmt.Printf("\nHere is the updated list on ice cream varieties %d", len(xj))
	for i := 0; i < len(xj); i++ {
		fmt.Printf("\n%d\t%#v", i+1, xj[i])
	}
	/*
		xk := xj
		xk[0] = "Rainbow Lovers"
		fmt.Printf("\nHere is the updated list on ice cream varieties %d", len(xk))
		for i := 0; i < len(xk); i++ {
			fmt.Printf("\n%d\t%#v\t\t\t%#v", i+1, xj[i], xk[i])
		}
	*/
	fmt.Println("\n*************************************")
	fmt.Println("\n updated icecreams from the menu")
	fmt.Println("\n*************************************")
	xk := make([]string, 0)
	xl := make([]string, 22)
	xk = append(xk[0:], xj[:]...)
	copy(xl, xj)
	xl[0] = "Rainbow lover1"
	xk[0] = "Rainbow lover2"
	fmt.Printf("\nHere is the updated list on ice cream varieties %d", len(xk))
	sort.Strings(xl)
	sort.Strings(xk)
	for i := 0; i < len(xk); i++ {
		fmt.Printf("\n%d\t%#v\t\t\t%#v", i+1, xj[i], xk[i])
		fmt.Printf("\n%d\t%#v\t\t\t%#v", i+1, xj[i], xl[i])
	}

}
