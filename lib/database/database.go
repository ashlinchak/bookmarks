package database

import (
	"os"

	"github.com/ashlinchak/bookmarks/lib/repositories"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Database holds a connection
type Database struct {
	Conn               *gorm.DB
	TagRepository      *repositories.TagRepository
	BookmarkRepository *repositories.BookmarkRepository
}

// GetDatabase creates connection to database
func GetDatabase() *Database {
	db := &Database{}

	db.Conn = createConnection()
	initRepositories(db)

	return db
}

func createConnection() *gorm.DB {
	dbPath := "data/bookmarks.db"

	if dbPathEnv := os.Getenv("BOOKMARKS_DB_PATH"); len(dbPathEnv) > 0 {
		dbPath = dbPathEnv
	}

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		panic("failed to connect database")
	}

	return db
}

func initRepositories(db *Database) {
	db.TagRepository = &repositories.TagRepository{Conn: db.Conn}
	db.BookmarkRepository = &repositories.BookmarkRepository{Conn: db.Conn}
}
