package Common

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func CreateTable(db *sql.DB) {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS testmodule (
	       id SERIAL PRIMARY KEY,
	       name TEXT
       )`)
	if err != nil {
		log.Println("Error Occured: ", err)
		return
	}
	log.Println("Table created")
}

func InsertIntoDB(db *sql.DB, table string, columns []string, data ...interface{}) {
	query := "INSERT INTO " + table + " ("
	for i, col := range columns {
		if i > 0 {
			query += ", "
		}
		query += col
	}
	query += ") VALUES ("
	for i := range data {
		if i > 0 {
			query += ", "
		}
		query += "$" + string('1'+i)
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

func SelectTestmodule(db *sql.DB, table string) (*sql.Rows, error) {
	rows, err := db.Query("SELECT * FROM " + table)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return rows, nil
}

func Connect(connStr string) *sql.DB {
	// Example connStr: "user=youruser password=yourpass dbname=yourdb sslmode=disable"
	database, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	return database
}
