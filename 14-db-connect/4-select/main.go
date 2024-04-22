package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := dbConnection()
	defer db.Close()
	selectQr(db)
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

type Team struct {
	name   string
	zone   string
	points int
}

func selectQr(db *sql.DB) {
	// result, err := db.Query("SELECT nombre, puntos FROM equipos WHERE zona = 'a'")
	result, err := db.Query("SELECT nombre, puntos FROM equipos")
	if err != nil {
		fmt.Println("error showing data")
		panic(err.Error())
	}
	for result.Next() {
		var team Team
		err = result.Scan(&team.name, &team.points)
		if err != nil {
			fmt.Println("error scanning values")
			panic(err.Error())
		}
		fmt.Println(team)
	}

}
