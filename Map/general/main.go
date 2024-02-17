package main

import "fmt"

func main() {
	ai := map[string]int{
		"Rams":  35,
		"Gayus": 32,
		"Ezhil": 3,
		"Veda":  7,
	}
	fmt.Printf("I am %d years old\n", ai["Rams"])
	fmt.Printf("My family strength is with %d people\n", len(ai)-1)
	fmt.Println(ai)
	for a, b := range ai {
		fmt.Printf("\nAge of %v \t: %d", a, b)
	}
	fmt.Println("\n* * My family strenght now * *")
	ai["Bhavani"] = 64
	ai["Azagusundaram"] = 76
	ai["Vanthru"] = 14
	ai["Dharshu"] = 20
	ai["Ratchu"] = 18
	//	fmt.Println(ai)
	for a, _ := range ai {
		fmt.Printf("\n %v", a)
	}
	delete(ai, "Rams")
	fmt.Println("\n* * Without me my family strenght now * *")
	for a, _ := range ai {
		fmt.Printf("\n %v", a)
	}
}
