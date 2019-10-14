package database

import (
	"fmt"

	"github.com/ashlinchak/bookmarks/lib/models"
)

// Migrate database
func (db *Database) Migrate() {
	fmt.Println("Start: migrate DB")

	db.Conn.AutoMigrate(&models.Bookmark{}, &models.Tag{})

	fmt.Println("Finished: migrate DB")
}
