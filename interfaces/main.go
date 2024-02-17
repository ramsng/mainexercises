package main

import "fmt"

type games struct {
	name    string
	players int
	buffers int
	coach   string
	kitreqd bool
}

type gamesperday int
type game interface {
	gameplan()
}

func intergameplan(g game) {
	fmt.Println("**** Print through interface ****")
	g.gameplan()
	var gd2 gamesperday
	gd2 = 24
	gd2.schedule()
}

func (g games) gameplan() {
	fmt.Printf("Game: %v;\nNumber of players: %v and buffers:%v;\nRequired Kit availablity: %v\n", g.name, g.players, g.buffers, g.kitreqd)
}

func (gd gamesperday) schedule() {
	fmt.Printf("Game planned per day: %v\n", gd)
}

/*
func teamplan(s schedule) {
	gameplan()
}
*/
func main() {
	cricket := games{
		name:    "Cricket",
		players: 11,
		buffers: 5,
		coach:   "Ravi",
		kitreqd: true,
	}
	vball := games{
		name:    "Volleyball",
		players: 6,
		buffers: 5,
		coach:   "Raju",
		kitreqd: false,
	}

	var gd1 gamesperday
	gd1 = 42
	gd1.schedule()
	//call through methods.
	cricket.gameplan()
	vball.gameplan()
	//call through interface
	intergameplan(cricket)
	intergameplan(vball)
}
