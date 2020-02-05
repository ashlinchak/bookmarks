package database

import (
	"github.com/ashlinchak/bookmarks/lib/models"
)

// Migrate database
func (db *Database) Migrate() {
	db.Conn.AutoMigrate(&models.Bookmark{}, &models.Tag{})
}
