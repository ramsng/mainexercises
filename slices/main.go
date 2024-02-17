// array concepts
package main

import "fmt"

func main() {
	/*
		var x [5]int
		x[3] = 42
		fmt.Println("Display array :", x)
		y := [5]int{42, 21, 42, 23, 35}
		fmt.Println("Display array 1 : ", y)
		u := [...]string{"Mango", "Banana", "Jackfruit"}
		fmt.Printf("%#v\t of type %#T\n", u, u)
		var w [2]int
		w[0] = 15
		w[1] = 20
		fmt.Println("Display array 2 : ", w)
		v := [...]string{"Mango", "Banana", "Jackfruit"}
		fmt.Printf("%#v\t of type %#T\n", v, v[1])
		fmt.Println(len(u))
		fmt.Println(len(x))
		fmt.Println(len(y))
	*/
	a := [...]string{
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
	fmt.Printf("Ice creams - Total varieties %#d\n", len(a))
	for i := 0; i < len(a); i++ {
		fmt.Printf("%d.\t%s\n", i+1, a[i])
	}
}
