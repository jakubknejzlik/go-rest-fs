package main

import (
	"./model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// NewDB ...
func NewDB(url string) (*gorm.DB, error) {
	// Migrate the schema
	db, err := gorm.Open("sqlite3", url)
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&model.File{})
	return db, nil
}
