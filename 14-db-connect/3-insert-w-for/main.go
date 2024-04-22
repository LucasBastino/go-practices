package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Team struct {
	name   string
	zone   string
	points int
}

type Player struct {
	name     string
	age      int
	position string
}

func main() {
	db := dbInit()
	defer db.Close()

	oroz := Player{"Oroz", 25, "MED"}
	vivaldo := Player{"Vivaldo", 39, "ARQ"}
	malcorra := Player{"Malcorra", 32, "MED"}
	campaz := Player{"Campaz", 24, "DEL"}
	chacarita := Team{"Chacarita Jrs", "A", 49}
	rosario := Team{"Rosario Central", "A", 42}
	defensores := Team{"Defensores de Belgrano", "B", 44}
	sanTelmo := Team{"San Telmo", "B", 40}
	teams := []Team{chacarita, rosario, defensores, sanTelmo}
	fmt.Println(teams)
	players := []Player{oroz, vivaldo, malcorra, campaz}
	// insertTeams(db, teams)
	insertPlayers(db, players)
}

func dbInit() *sql.DB {
	host := "w0x.h.filess.io"
	user := "goProjects_cowboydish"
	password := "databasefilessio"
	dbName := "goProjects_cowboydish"
	port := "3307"

	connString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, dbName)
	db, err := sql.Open("mysql", connString)
	if err != nil {
		fmt.Println("error connecting with mysql database")
		panic(err.Error())
	}
	fmt.Println("database connection ok")

	err = db.Ping()
	if err != nil {
		fmt.Println("error with db.Ping")
		panic(err.Error())
	}
	return db
}

// func insertTeams(db *sql.DB, teams []Team) {
// 	queryString := "INSERT INTO Teams (nombre, zona, puntos) VALUES "
// 	for _, e := range teams {
// 		// se le agregan los values
// 		queryString += fmt.Sprintf("('%s', '%s', '%v'),", e.name, e.zone, e.points)
// 	}

// 	//  elimino la ',' al final del string
// 	queryString = queryString[:len(queryString)-1]

// 	insert, err := db.Query(queryString)
// 	if err != nil {
// 		fmt.Println("error inserting data")
// 		panic(err.Error())
// 	}
// 	defer insert.Close()

// }

func insertPlayers(db *sql.DB, players []Player) {
	insertPlayersQueryString := "INSERT INTO jugadores (nombre, edad, posicion) VALUES "
	for _, player := range players {
		insertPlayersQueryString += fmt.Sprintf("('%s', '%v', '%s'),", player.name, player.age, player.position)
	}
	insertPlayersQueryString = insertPlayersQueryString[:len(insertPlayersQueryString)-1]
	insert, err := db.Query(insertPlayersQueryString)
	if err != nil {
		fmt.Println("error inserting players")
		panic(err.Error())
	}
	defer insert.Close()
}
