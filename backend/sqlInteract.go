package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	//initEventTables()
	ParseRSS()
}

func initEventTables() {

	dbPath := "./backend/owl_gather.db"
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	createTables(db)

	/*tables, err := db.Query("SELECT * FROM sqlite_master")
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

	fmt.Printf("End statement") */

}

func createTables(db *sql.DB) {
	query := `CREATE TABLE IF NOT EXISTS EVENTS (
		EventId SERIAL PRIMARY KEY,
		Title VARCHAR(100) NOT NULL,
		StartDate DATE NOT NULL
	)`

	query1 := `CREATE TABLE IF NOT EXISTS EVENTSCATEGORY (
		EventId INT PRIMARY KEY,
		CategoryID INT NOT NULL
	)`

	query2 := `CREATE TABLE IF NOT EXISTS CATEGORIES (
		CategoryID SERIAL PRIMARY KEY,
		CategoryName VARCHAR(100) NOT NULL
	)`

	_, err := db.Exec(query)
	if err != nil {
		panic(err)
	}
	_, err = db.Exec(query1)
	if err != nil {
		panic(err)
	}
	_, err = db.Exec(query2)
	if err != nil {
		panic(err)
	}
}

/*func insertEvent(db *sql.DB, event Event) int {

}*/
