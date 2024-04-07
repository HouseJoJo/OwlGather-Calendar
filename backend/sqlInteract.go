package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	dbPath := "owl_gather.db"
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	fmt.Printf("after db??")

	tables, err := db.Query("SELECT * FROM sqlite_master")
	if err != nil {
		panic(err)
	}
	defer tables.Close()

	for tables.Next() {
		fmt.Printf("Entered .next")
		var tableName string
		err := tables.Scan(&tableName)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Table: %s\n", tableName)
	}

	fmt.Printf("End statement")

}
