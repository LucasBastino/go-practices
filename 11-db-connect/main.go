package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// db, err := sql.Open("mysql", "w0x.h.filess.io:databasefilessio@tcp(localhost:3307)/goProjects_cowboydish")
	db, err := sql.Open("mysql", "goProjects_cowboydish:databasefilessio@tcp(w0x.h.filess.io:3307)/goProjects_cowboydish")
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

	insert, err := db.Query("INSERT INTO equipos (nombre, zona, puntos) VALUES ('Chacarita', 'B', 47), ('Almirante', 'A', 41 );")
	if err != nil {
		fmt.Println("error inserting data in DB")
		panic(err.Error())
	}
	defer insert.Close()
	fmt.Println("Insert data has been run succesfully")

}
