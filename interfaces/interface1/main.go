package main

import "fmt"

type xface1 interface {
	interview()
}

type actor struct {
	firstnm string
	lastnm  string
	popular bool
}

func (a1 actor) interview() string {
	return fmt.Sprintf("Actor Name : %v %v; Popularity: %v\n", a1.firstnm, a1.lastnm, a1.popular)
}

func biodata(x xface1) {
	x.interview()
}

func main() {

	actor1 := actor{
		firstnm: "Christian",
		lastnm:  "Bale",
		popular: true,
	}
	/*actor2 := actor{
		firstnm: "Brad",
		lastnm:  "Pitt",
		popular: true,
	}
	*/
	actor3 := actor{
		firstnm: "Ryan",
		lastnm:  "Reynolds",
		popular: false,
	}
	//
	fmt.Println(actor1.interview())
	//fmt.Println(actor3.interview())
	var l string
	l = biodata(actor3)
	//fmt.Println(biodata(actor3))

}
