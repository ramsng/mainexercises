package main

import "fmt"

func main() {
	// functions: without arguments
	// functions: with one argument
	// functions: with multiple arguments
	// functions: with multiple arguments & multiple return
	input := findfruits()
	title, output := checkseedless("FRUIT", input)
	fmt.Println(title)
	for a, b := range output {
		fmt.Printf("%d . %s\n", a+1, b)
	}

}

func findfruits() []string {
	fruits := []string{"Banana", "Guava", "Pineapple", "Mango", "Strawberry"}
	return fruits
}

func checkseedless(flag string, flist []string) (string, []string) {
	fruitmapa := map[string]string{
		"Banana":     "YES",
		"Guava":      "YES",
		"Pineapple":  "NO",
		"Strawberry": "NO",
	}
	xa := []string{}
	if flag == "FRUIT" {
		for _, a := range flist {
			if _, ok := fruitmapa[a]; ok {
				if fruitmapa[a] == "YES" {
					xa = append(xa, a)
				} else {
					continue
				}
				//		fmt.Println(a, v)
			} else {
				continue
			}
		}
	}
	return fmt.Sprintf("SEEDLESS %s LIST", flag), xa
}
func init() {
	fmt.Println("** Welcome to fruit bar **")
}
