package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// NewDB ...
func NewDB(url string) (*gorm.DB, error) {
	return gorm.Open("sqlite3", "test.db")
}
