package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func test() {

	dbPath := "./backend/owl_gather.db"
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

	tableColumns, err := tables.Columns()
	if err != nil {
		panic(err)
	}
	for i, col := range tableColumns {
		fmt.Println(i, col)
	}

	for tables.Next() {
		fmt.Printf("Entered .next")
		var tableName string
		err := tables.Scan(&tableName) //causes error. expected 5 args not 1?
		if err != nil {
			panic(err)
		}

		fmt.Printf("Table: %s\n", tableName)
	}

	fmt.Printf("End statement")

}
