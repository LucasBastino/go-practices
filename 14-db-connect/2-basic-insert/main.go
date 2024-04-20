package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	host := "w0x.h.filess.io"
	user := "goProjects_cowboydish"
	password := "databasefilessio"
	dbName := "goProjects_cowboydish"
	port := "3307"
	connString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, dbName)
	// db, err := sql.Open("mysql", "goProjects_cowboydish:databasefilessio@tcp(w0x.h.filess.io:3307)/goProjects_cowboydish")
	db, err := sql.Open("mysql", connString)
	if err != nil {
		fmt.Println("error validating sql.Open arguments")
		panic(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("error verifying connection with db.Ping")
		panic(err.Error())
	}
	fmt.Println("Successfully connection to DB!")

	tabla := "equipos"
	equipo := "Chacarita"
	zona := "A"
	puntos := 45
	// insert, err := db.Query("INSERT INTO equipos (nombre, zona, puntos) VALUES ('Chacarita', 'B', 47), ('Almirante', 'A', 41 );")
	insert, err := db.Query(fmt.Sprintf("INSERT INTO %s (nombre, zona, puntos) VALUES ('%s', '%s', '%v')", tabla, equipo, zona, puntos))
	if err != nil {
		fmt.Println("error inserting data in DB")
		panic(err.Error())
	}
	defer insert.Close()
	fmt.Println("Insert data has been run succesfully")

}
