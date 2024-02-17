package main

import (
	"fmt"
	"sort"
)

// custom sort
type User struct {
	Name string
	Age  int
	Auth bool
}
type ByName []User

func (b ByName) Len() int {
	return len(b)
}
func (b ByName) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}
func (b ByName) Less(i, j int) bool {
	return b[i].Name < b[j].Name
}
func main() {
	user1 := ByName{
		User{Name: "Rams",
			Age:  36,
			Auth: true},
		User{"Ravi", 35, true},
		User{"Raja", 39, true},
		User{"Ramesh", 40, true},
	}
	fmt.Println(user1)
	sort.Sort(user1)
	fmt.Println(user1)
}
