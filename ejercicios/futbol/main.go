package main

import (
	"fmt"
	"math"
	"math/rand"

	"github.com/LucasBastino/practicas-go/sliceMethods"
)

type League struct {
	teams []Team
}

type Team struct {
	name      string
	fundation int
	budget    float64
	players   []Player
	points    int
}

type Player struct {
	name       string
	age        int
	position   string
	valoration float64
	salary     float64
	value      float64
	injured    bool
}

func main() {
	fmt.Println("Futbol")
	oroz := Player{"Oroz", 25, "MED", 52, 5000, 2, false}
	vivaldo := Player{"Vivaldo", 39, "ARQ", 52, 3000, 1.5, true}
	malcorra := Player{"Malcorra", 32, "MED", 88, 6000, 3, false}
	campaz := Player{"Campaz", 24, "DEL", 86, 5500, 2.5, false}
	chacarita := Team{"Chacarita Jrs", 1906, 20, []Player{oroz, vivaldo}, 25}
	central := Team{"Rosario Central", 1889, 30, []Player{campaz, malcorra}, 22}
	league := League{[]Team{chacarita, central}}
	fmt.Println(chacarita)
	fmt.Println(central)
	chacarita.buyPlayer("Malcorra", &central, 2)
	// chacarita.buyPlayer("Campaz", &central, 2)
	// central.buyPlayer("Oroz", &chacarita, 1.5)
	chacarita.showInfo()
	central.showInfo()
	// chacarita.teamMedia()
	// central.teamMedia()
	league.playMatch(chacarita, central)
}

// func (e *Team) eliminarPlayer(field string, value any) {
// 	e.players = sliceMethods.DeleteOneByField(e.players, field, value)
// 	// fmt.Println(e.players)
// }

func (e *Team) buyPlayer(Player string, sellerTeam *Team, valor float64) {
	i := sliceMethods.FindIndexByField(sellerTeam.players, "name", Player)
	e.players = append(e.players, sellerTeam.players[i])
	e.budget -= valor
	sellerTeam.players = sliceMethods.DeleteOneByIndex(sellerTeam.players, i)
	sellerTeam.budget += valor
}

func (e Team) showInfo() {
	fmt.Println(e.name)
	fmt.Println("Budget:", e.budget)
	for _, Player := range e.players {
		fmt.Println(Player.name)
	}
}

func (e Team) teamMedia() float64 {
	var totalValue float64
	for _, value := range e.players {
		totalValue += value.valoration
	}
	media := totalValue / float64(len(e.players))
	fmt.Println("The media of", e.name, "is", totalValue/float64(len(e.players)))
	return media
}

func (l League) playMatch(homeTeam Team, awayTeam Team) {
	homeGoals := math.Round(rand.Float64() * (homeTeam.teamMedia()) * 1.40 / 35)
	awayGoals := math.Round(rand.Float64() * (awayTeam.teamMedia()) * 1.40 / 35)
	fmt.Printf("%s %v-%v %s", homeTeam.name, homeGoals, awayGoals, awayTeam.name)
}
