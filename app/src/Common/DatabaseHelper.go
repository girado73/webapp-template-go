package Common

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func CreateDb(db *sql.DB) {

	_, err := db.Exec("CREATE DATABASE IF NOT EXISTS testmodule")
	if err != nil {
		log.Println("Error Occured: ", err)
		return
	}
	log.Println("Database created")
}

func CreateTable(db *sql.DB) {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS testmodule (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT)")
	if err != nil {
		log.Println("Error Occured: ", err)
		return
	}
	log.Println("Table created")
}

func InsertIntoDB(db *sql.DB, table string, data ...interface{}) {
	query := "INSERT INTO " + table + " VALUES ("
	for i := range data {
		if i > 0 {
			query += ", "
		}
		query += "?"
	}
	query += ")"

	stmt, err := db.Prepare(query)
	if err != nil {
		log.Println("Error preparing statement:", err)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(data...)
	if err != nil {
		log.Println("Error executing statement:", err)
	}
	log.Println("Inserted into database", data)
}

func selectTestmodule(db *sql.DB, table string) *sql.Rows {

	rows, err := db.Query("SELECT * FROM " + table)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	return rows
}

func Connect(dbFile string) *sql.DB {
	// Open the database
	database, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		panic(err)
	}
	return database
}
