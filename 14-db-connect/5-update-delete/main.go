package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := dbConnection()
	defer db.Close()
	buyPlayer(db, "Malcorra", 2, 1)
}

func dbConnection() *sql.DB {
	host := "w0x.h.filess.io"
	user := "goProjects_cowboydish"
	password := "databasefilessio"
	dbName := "goProjects_cowboydish"
	port := "3307"

	connString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, dbName)
	db, err := sql.Open("mysql", connString)
	if err != nil {
		fmt.Println("error connecting to database")
		panic(err.Error())
	}
	return db
}

func buyPlayer(db *sql.DB, playerName string, idSellerTeam int, idBuyerTeam int) {
	findPlayer(db, playerName, idSellerTeam)
	// insertPlayer(db, playerName, id)
	// deletePlayer(db, playerName, idSellerTeam)
}

func findPlayer(db *sql.DB, playerName string, idSellerTeam int) {
	queryString := fmt.Sprintf("SELECT player.name, player.age, team.name FROM player INNER JOIN team ON player.idTeam = team.idTeam WHERE player.name = '%s'", playerName)
	// queryString := "SELECT * FROM player"
	findSelect, err := db.Query(queryString)
	if err != nil {
		fmt.Println("error in select query")
		panic(err.Error())
	}
	for findSelect.Next() {
		var show ShowPlayer
		err := findSelect.Scan(&show.name, &show.age, &show.teamName)
		if err != nil {
			fmt.Println("error scanning data")
			panic(err.Error())
		}
		fmt.Println(show)
	}
	defer findSelect.Close()
}

type ShowPlayer struct {
	idPlayer int
	name     string
	age      int
	teamName string
}
