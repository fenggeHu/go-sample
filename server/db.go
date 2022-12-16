package server

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"time"

	// dialect for sqlite
	_ "gorm.io/driver/sqlite"
	// sqlite driver
	//_ "github.com/mattn/go-sqlite3"
)

func init() {
	log.Printf("db init: %s", time.Now())
}

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
		&Config{},
		&Movie{},
		&Subtitle{},
	)
}
