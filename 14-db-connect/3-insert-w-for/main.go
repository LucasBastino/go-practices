package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Equipo struct {
	nombre string
	zona   string
	puntos int
}

func main() {
	db := dbInit()
	defer db.Close()

	chacarita := Equipo{"Chacarita Jrs", "A", 49}
	rosario := Equipo{"Rosario Central", "A", 42}
	defensores := Equipo{"Defensores de Belgrano", "B", 44}
	sanTelmo := Equipo{"San Telmo", "B", 40}
	equipos := []Equipo{chacarita, rosario, defensores, sanTelmo}
	insertDB(db, equipos)

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

func insertDB(db *sql.DB, equipos []Equipo) {
	queryString := "INSERT INTO equipos (nombre, zona, puntos) VALUES "
	for _, e := range equipos {
		// se le agregan los values
		queryString += fmt.Sprintf("('%s', '%s', '%v'),", e.nombre, e.zona, e.puntos)
	}

	//  elimino la ',' al final del string
	queryString = queryString[:len(queryString)-1]

	insert, err := db.Query(queryString)
	if err != nil {
		fmt.Println("error inserting data")
		panic(err.Error())
	}
	defer insert.Close()

}
