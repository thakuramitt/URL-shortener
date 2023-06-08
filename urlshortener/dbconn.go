package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func dbcon() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("db/urls.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	db.AutoMigrate(&ShortURL{})
	return db, nil
}
