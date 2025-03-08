package main

import (
	"webapp-template-go/src/Common"
)

func main() {

	// Open the database connection

	db := Common.Connect("./database.db")

	// Create the database
	//Common.CreateDb(db)
	Common.CreateTable(db)
	Common.InsertIntoDB(db, "testmodule", nil, "testmodule1")
}
