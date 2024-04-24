package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := dbConnection()
	defer db.Close()
	buyPlayer(db, "Vivaldo", 1, 2)
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
	queryString := fmt.Sprintf("UPDATE player SET idTeam = '%v' WHERE name = '%s' AND idTeam = '%v'", idBuyerTeam, playerName, idSellerTeam)
	updatePlayer, err := db.Query(queryString)
	if err != nil {
		fmt.Println("error in UPDATE query")
		panic(err.Error())
	}
	defer updatePlayer.Close()
}
