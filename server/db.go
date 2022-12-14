package server

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	// dialect for sqlite
	_ "gorm.io/driver/sqlite"
	// sqlite driver
	//_ "github.com/mattn/go-sqlite3"
)

// OpenDB will create new database connection to Sqlite
func OpenDB(dbPath string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(dbPath))
	if err != nil {
		return nil, err
	}

	return db, nil
}

// Migrate will do migration of models
func Migrate(db *gorm.DB) {
	db.AutoMigrate(
		&Movie{},
		&Subtitle{},
	)
}
