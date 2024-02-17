package main

import "fmt"

type dog struct {
	name     string
	category string
}

// pointer sematic
func (d *dog) walk() {
	fmt.Printf("I am %v and am %v, so can walk\n", d.name, d.category)
}

func (d *dog) run() {
	d.name = "Roger"
	d.category = "Young"
	fmt.Printf("I am %v and am %v, so can run\n", d.name, d.category)
}

type move interface {
	walk()
	run()
}

func youngin(m move) {
	m.walk()
	m.run()
}

func main() {
	name := "Stefi"
	category := "Young"
	dog1 := dog{
		name:     name,
		category: category,
	}
	//dog2 := dog{"Mike", "Old"}
	//dog3 := dog{"Jimmy", "Young"}
	youngin(&dog1)
	dog1.walk()
	dog1.run()

	//	youngin(&dog2)
	//	youngin(&dog3)
}

// pointer sematic
/*
func (d *dog) walk() {
	fmt.Printf("I am %v and am %v, so can walk\n", d.name, d.category)
}

func (d *dog) run() {
	d.name = "Roger"
	d.category = "young"
	fmt.Printf("I am %v and am %v, so can run\n", d.name, d.category)
}

type move interface {
	walk()
	run()
}

func youngin(m move) {
	m.walk()
	m.run()
}

func main() {
	dog1 := dog{"Stefi", "Young"}
	dog2 := dog{"Mike", "Old"}
	dog3 := dog{"Jimmy", "Young"}
	youngin(&dog1)
	youngin(&dog2)
	youngin(&dog3)
}

// Value sematic
/*
func (d dog) walk() {
	fmt.Printf("I am %v and am %v, so can walk\n", d.name, d.category)
}

func (d dog) run() {
	d.name = "Roger"
	d.category = "young"
	fmt.Printf("I am %v and am %v, so can run\n", d.name, d.category)
}

type move interface {
	walk()
	run()
}

func youngin(m move) {
	m.walk()
	m.run()
}

func main() {
	dog1 := dog{"Stefi", "Young"}
	dog2 := dog{"Mike", "Old"}
	dog3 := dog{"Jimmy", "Young"}
	youngin(dog1)
	youngin(dog2)
	youngin(dog3)
}
*/
