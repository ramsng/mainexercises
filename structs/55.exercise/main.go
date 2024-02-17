package main

import "fmt"

func main() {
	type engine struct {
		electric bool
	}
	type car struct {
		eng     engine
		carmake string
		model   string
		doors   int
		color   string
	}
	car1 := car{
		eng:     engine{true},
		carmake: "Lexus",
		model:   "SUV",
		doors:   4,
		color:   "White",
	}
	car2 := car{
		eng: engine{
			electric: false},
		carmake: "Toyota",
		model:   "Coupe",
		doors:   2,
		color:   "Black",
	}
	fmt.Println(" ** 2024 - Jan Car List ** ")
	fmt.Println(" Car make : ", car1.carmake)
	fmt.Println("    Model : ", car1.model)
	fmt.Println("    Doors : ", car1.doors)
	fmt.Println("    Color : ", car1.color)
	fmt.Println("    Type  : ", car1.eng)
	if car1.eng.electric == true {
		fmt.Println(" Electric start : YES")
	} else {
		fmt.Println(" Electric start : NO")
	}
	fmt.Println(" Car make : ", car2.carmake)
	fmt.Println("    Model : ", car2.model)
	fmt.Println("    Doors : ", car2.doors)
	fmt.Println("    Color : ", car2.color)
	fmt.Println("    Type  : ", car2.eng.electric)
	if car2.eng.electric == true {
		fmt.Println(" Electric start : YES")
	} else {
		fmt.Println(" Electric start : NO")
	}

}
