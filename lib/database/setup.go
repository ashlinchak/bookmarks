package database

// Setup database.
func (db *Database) Setup() {
	db.Migrate()
}
