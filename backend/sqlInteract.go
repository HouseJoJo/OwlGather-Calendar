package main

import (
	"database/sql"
	"fmt"
	"regexp"
	"strings"

	_ "github.com/mattn/go-sqlite3"

	"github.com/mmcdole/gofeed"
)

type Event struct {
	StartDate  string
	EventTitle string
	EventID    int
	Categories []string
}

type Categories struct {
	category string
	catID    int
}

var Events []Event
var UniqueCategories []Categories

func main() {
	//initEventTables()
	initEventTables()
	ParseRSS()
	//printEvents(Events)
	UniqueCategories = getUniqueCategories(Events, UniqueCategories) //at this point Events should be filled with structs and empty ID, UniqueCategories should be filled with Categories structs with no ID as well

	//TODO: insert UniqueCategories into CATEGORIES table and fill slice with IDS
	//TODO: insert Events into EVENTS table and fill slice with EventIDs
	//TODO: insert Events and UniqueCategories into EVENT_CATEGORIES where every category/Event pair in Events slice is a new entry (catID/eventID)

	dbPath := "./backend/owl_gather.db"
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	fillEventTable(Events, db)
	fmt.Println("Events filled.")

	fillCategoriesTable(UniqueCategories, db)
	fmt.Println("Categories filled.")

	fillEventCatJoinTable(UniqueCategories, Events, db)
	fmt.Println("Event-Cat Join Table filled.")

}

func fillEventCatJoinTable(uniqueCategories []Categories, events []Event, db *sql.DB) {
	for _, event := range events {
		for _, category := range event.Categories {
			catId := -1
			for i := range uniqueCategories {
				if category == uniqueCategories[i].category {
					catId = uniqueCategories[i].catID
				}
			}
			res, err := db.Exec("INSERT INTO EVENT_CATEGORIES(EventId, CategoryID) VALUES(?,?)", event.EventID, catId)
			if err != nil {
				fmt.Println(err)
				fmt.Println(res)
			}
		}
	}
}

func fillCategoriesTable(categories []Categories, db *sql.DB) {
	for i := range categories {
		cId, err := InsertCategoriesToDB(db, categories[i].category)
		if err != nil {
			fmt.Println(err)
		}
		categories[i].catID = cId
	}
}

func fillEventTable(events []Event, db *sql.DB) {
	for i := range events {
		eId, err := InsertEventToDB(db, events[i])
		if err != nil {
			fmt.Println(err)
		}
		events[i].EventID = eId
	}
}

func getUniqueCategories(events []Event, categories []Categories) (out []Categories) { //method that takes in slices of Event structs and return a string slice with only the unique values of categories, intended to be a helper function to fill CATEGORIES table of DB.
	mapElements := make(map[string]bool)

	for _, event := range events {
		for _, category := range event.Categories {
			if _, inMap := mapElements[category]; !inMap {
				mapElements[category] = true
				categories = append(categories, Categories{category, -1})
			}
		}
	}
	return categories
}

func InsertEventToDB(db *sql.DB, event Event) (int, error) { //Events struct should be iterated over and used to call this function to enter all items to EVENT DB
	res, err := db.Exec("INSERT INTO EVENTS(Title, StartDate) VALUES(?,?);", event.EventTitle, event.StartDate)
	if err != nil {
		return -1, err //Return int -1 and error code
	}

	var id int64
	if id, err = res.LastInsertId(); err != nil {
		return -1, err
	}
	return int(id), nil //will return id to method which should be used to update original Events Struct

}

func InsertCategoriesToDB(db *sql.DB, category string) (int, error) {
	res, err := db.Exec("INSERT INTO CATEGORIES (CategoryName) VALUES(?);", category) //CategoryID, CategoryName
	if err != nil {
		return -1, err
	}

	var id int64
	if id, err = res.LastInsertId(); err != nil {
		return -1, err
	}
	return int(id), nil
}

func printEvents(slice []Event) {
	for i := range slice {
		fmt.Println(slice[i])
		fmt.Println(slice[i].EventID)
	}
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
		EventId INTEGER PRIMARY KEY,
		Title VARCHAR(100) NOT NULL,
		StartDate DATE NOT NULL
	)`

	query1 := `CREATE TABLE IF NOT EXISTS EVENT_CATEGORIES (
		EventId INT NOT NULL,
		CategoryID INT NOT NULL,
		PRIMARY KEY (EventId, CategoryID)
	)`

	query2 := `CREATE TABLE IF NOT EXISTS CATEGORIES (
		CategoryID INTEGER PRIMARY KEY,
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

func ParseRSS() {
	fp := gofeed.NewParser()

	feed, err := fp.ParseURL("https://owllife.kennesaw.edu/events.rss")
	if err != nil {
		fmt.Println("Error parsing RSS feed:", err)
		return
	}

	if len(feed.Items) == 0 {
		fmt.Println("No items found in the RSS feed.")
		return
	}

	fmt.Println(feed.String())
	var mar, ken, other, online int = 0, 0, 0, 0

	for _, item := range feed.Items {
		//fmt.Println("-------------------------")
		currTitle := item.Title
		//fmt.Println("\nTitle: ", item.Title)

		//fmt.Println("\nLink: ", item.Link)
		/*	fmt.Println("\nAuthor ", item.Author)
			fmt.Println("\nImage ", item.Image)*/
		var currCategories []string
		for _, category := range item.Categories {
			//fmt.Printf("Categ. %d: %s\n", index+1, category)
			currCategories = append(currCategories, category)
		}
		//fmt.Println("\nCategories", item.Categories)
		//fmt.Println("\nLocation: ", item.Custom["location"])

		fmt.Println("\nDesc. ", item.Description)
		//fmt.Println("\nDesc: ", parseDesc(item.Description))
		//fmt.Println("\nDateTime: ", parseDate(item.Description))
		currDateTime := parseDate(item.Description)
		//fmt.Println("\nLocation. ", parseStrings(item.Custom, `location:`, `start`))
		Events = append(Events, Event{currDateTime, currTitle, -1, currCategories})
		//fmt.Println("Desc:", item.Description)
		//campus := ("\nModality: " + modality(item.Categories, item.Custom["location"]))
		if modality(item.Categories, item.Custom["location"]) == `Kennesaw` {
			ken++
		}
		if modality(item.Categories, item.Custom["location"]) == `Online` {
			online++
		}
		if modality(item.Categories, item.Custom["location"]) == `Marietta` {
			mar++
		}
		if modality(item.Categories, item.Custom["location"]) == `Other` {
			other++
		}

		//fmt.Println(campus)

	}

	fmt.Printf("Mar. %d, Ken. %d, Onl. %d, Other %d", mar, ken, online, other)

}

func modality(categories []string, location string) (out string) {
	if strings.Contains(location, `.com`) {
		return `Online`
	}
	if strings.Contains(location, `Marietta`) {
		return `Marietta`
	}
	if strings.Contains(location, `Kennesaw`) {
		return `Kennesaw`
	}

	for _, category := range categories {
		if strings.Contains(category, `Kennesaw`) {
			return `Kennesaw`
		}
		if strings.Contains(category, `Marietta`) {
			return `Marietta`
		}
	}
	return `Other`
}

func parseDesc(inString string) (out string) {
	startIndex := strings.Index(inString, `<div class="p-description description"><p>`)
	if startIndex == -1 {
		return "N/A"
	}

	endIndex := startIndex + strings.Index(inString[startIndex:], "</div>")
	if endIndex == -1 {
		return "N/A"
	}

	strippedString := inString[startIndex:endIndex]
	re := regexp.MustCompile(`<.*?>`)
	return re.ReplaceAllString(strippedString, "")
}

func parseStrings(inString, before, after string) (out string) {
	startIndex := strings.Index(inString, before)
	if startIndex == -1 {
		return
	}

	endIndex := startIndex + strings.Index(inString[startIndex:], after)
	if endIndex == -1 {
		return
	}

	return inString[(startIndex + len(before)):endIndex]
}

func parseDate(inString string) (out string) {
	startIndex := strings.Index(inString, `<time class="dt-start dtstart"`)
	if startIndex == -1 {
		return "N/A"
	}

	endIndex := startIndex + strings.Index(inString[startIndex:], `</time>`)
	if endIndex == -1 {
		return "N/A"
	}

	strippedString := inString[startIndex:endIndex]
	re := regexp.MustCompile(`<.*?>`)
	return re.ReplaceAllString(strippedString, "")
}
