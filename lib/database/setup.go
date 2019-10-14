package database

import "fmt"

// Setup database.
func (db *Database) Setup() {
	fmt.Println("Start: setup DB")

	db.Migrate()

	fmt.Println("Finished: setup DB")
}
