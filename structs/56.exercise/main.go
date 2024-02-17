package main

import "fmt"

func main() {
	mapa := map[string]int{}
	mapa["Rams"] = 4
	type party struct {
		first   string
		friends map[string]int
		drinks  []string
	}
	friend1 := party{
		first: "Venkat",
		friends: map[string]int{
			"Rams": 4,
		},
		drinks: []string{"Jameson", "Bombay Sapphire"},
	}
	friend2 := struct {
		first   string
		friends map[string]int
		drinks  []string
	}{
		first: "Sudharson",
		friends: map[string]int{
			"Gokul": 4,
		},
		drinks: []string{"Jack Daniels", "Voodoo"},
	}
	fmt.Println(friend1.first, " friend is ", friend1.friends, "Total Liquor reqd :", friend1.drinks)
	fmt.Println(friend2.first, " friend is ", friend2.friends, "Total Liquor reqd :", friend2.drinks)
	fmt.Printf("%v friend is %v ; Total Liquor reqd : %v", friend2.first, friend2.friends, friend2.drinks)
}
