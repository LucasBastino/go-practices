package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := dbConnection()
	defer db.Close()
	buyPlayer(db)
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

func buyPlayer(db *sql.DB) {
	equipoParaBorrar := "equipo1"
	delete, err := db.Query(fmt.Sprintf("DELETE FROM equipos WHERE nombre = '%v'", equipoParaBorrar))
	if err != nil {
		fmt.Println("error in first query")
		panic(err.Error())
	}
	defer delete.Close()
	insert2, err := db.Query("INSERT INTO equipos (nombre, zona, puntos) VALUES ('equipo3', 'd', '4')")
	if err != nil {
		fmt.Println("error in second insert")
		panic(err.Error())
	}
	defer insert2.Close()
	insert3, err := db.Query("INSERT INTO equipos (nombre, zona, puntos) VALUES ('equipo4', 'd', '4')")
	if err != nil {
		fmt.Println("error in second insert")
		panic(err.Error())
	}
	defer insert3.Close()
}
